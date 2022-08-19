package main

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/csmarchbanks/stravaql/strava"
	"github.com/csmarchbanks/stravaql/strava/activities"
	"github.com/csmarchbanks/stravaql/strava/model"
	"github.com/go-kit/log"
	"github.com/go-openapi/runtime"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/timestamp"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb/chunkenc"
)

type querier struct {
	logger        log.Logger
	ctx           context.Context
	client        *strava.StravaAPIV3
	auth          runtime.ClientAuthInfoWriter
	activityCache *activitySummaryCache
	mint, maxt    int64
}

func newQueryable(log log.Logger, client *strava.StravaAPIV3, auth runtime.ClientAuthInfoWriter, activityCache *activitySummaryCache) storage.Queryable {
	return storage.QueryableFunc(func(ctx context.Context, mint, maxt int64) (storage.Querier, error) {
		return &querier{
			ctx:           ctx,
			client:        client,
			auth:          auth,
			activityCache: activityCache,
			mint:          mint,
			maxt:          maxt,
		}, nil
	})
}

func (q *querier) Select(sortSeries bool, hints *storage.SelectHints, matchers ...*labels.Matcher) storage.SeriesSet {
	ctx := q.ctx

	metricName := metricName(matchers...)
	if metricName == "" {
		return storage.ErrSeriesSet(fmt.Errorf("metric name is required"))
	}

	interval := hints.Step
	if interval == 0 {
		interval = 60_000
	}

	switch metricName {
	case "strava_activities_total":
		allActivities, fetchStart := q.activityCache.Get()
		pageSize := int64(100)
		page := int64(1)
		//before := q.maxt / 1000
		after := fetchStart.Unix()
		params := activities.NewGetLoggedInAthleteActivitiesParamsWithContext(ctx).
			WithDefaults().
			WithPerPage(&pageSize).
			WithPage(&page).
			//	WithBefore(&before).
			WithAfter(&after)

		for {
			activities, err := q.client.Activities.GetLoggedInAthleteActivities(params, q.auth)
			if err != nil {
				return storage.ErrSeriesSet(err)
			}
			if len(activities.Payload) == 0 {
				break
			}
			allActivities = append(allActivities, activities.Payload...)
			page++
		}
		sort.Slice(allActivities, func(i, j int) bool {
			a := time.Time(allActivities[i].StartDate)
			b := time.Time(allActivities[j].StartDate)
			return a.Before(b)
		})
		q.activityCache.Put(allActivities)
		return newActivityCountSeriesSet(allActivities, q.mint, q.maxt, interval)
	case "strava_activity_moving_duration_seconds_total":
		return nil
	case "strava_activity_elapsed_duration_seconds_total":
		return nil
	default:
		return nil
	}
}

func (q *querier) LabelValues(name string, matchers ...*labels.Matcher) ([]string, storage.Warnings, error) {
	return nil, nil, nil
}

func (q *querier) LabelNames(matchers ...*labels.Matcher) ([]string, storage.Warnings, error) {
	return nil, nil, nil
}

// Close releases the resources of the Querier.
func (q *querier) Close() error {
	return nil
}

func metricName(matchers ...*labels.Matcher) string {
	for _, matcher := range matchers {
		if matcher.Name == "__name__" && matcher.Type == labels.MatchEqual {
			return matcher.Value
		}
	}
	return ""
}

type activityCountSeriesSet struct {
	activities []*model.SummaryActivity
	start, end time.Time
	interval   int64
	done       bool
}

func newActivityCountSeriesSet(activities []*model.SummaryActivity, start, end, interval int64) storage.SeriesSet {
	return &activityCountSeriesSet{
		activities: activities,
		start:      timestamp.Time(start),
		end:        timestamp.Time(end),
		interval:   interval,
	}
}

func (ss *activityCountSeriesSet) Next() bool {
	if ss.done {
		return false
	}
	ss.done = true
	return true
}

func (ss *activityCountSeriesSet) At() storage.Series {
	/*
		// The starting count is all rides before the start of the query.
		activityIndex := 0
		timestamps := []int64{}
		values := []float64{}
		for t := ss.start.UnixMilli(); t < ss.end.UnixMilli(); t += ss.interval {
			timestamps = append(timestamps, t)

			for _, activity := range ss.activities[activityIndex:] {
				if time.Time(activity.StartDate).Before(timestamp.Time(t)) {
					activityIndex++
				} else {
					break
				}
			}

			values = append(values, float64(activityIndex))
		}
	*/

	return &activityCountSeries{
		activities: ss.activities,
		start:      ss.start,
		end:        ss.end,
		interval:   ss.interval,
	}
}

func (ss *activityCountSeriesSet) Err() error {
	return nil
}

func (ss *activityCountSeriesSet) Warnings() storage.Warnings {
	return nil
}

type activityCountSeries struct {
	activities []*model.SummaryActivity
	start, end time.Time
	interval   int64
}

func (s *activityCountSeries) Labels() labels.Labels {
	return labels.FromStrings("__name__", "strava_activities_total")
}

func (s *activityCountSeries) Iterator() chunkenc.Iterator {
	return &activityCountIterator{
		activities: s.activities,
		start:      s.start,
		end:        s.end,
		interval:   s.interval,
		t:          timestamp.FromTime(s.start),
	}
}

type activityCountIterator struct {
	activities []*model.SummaryActivity
	start, end time.Time
	interval   int64
	t          int64
}

func (it *activityCountIterator) Seek(t int64) bool {
	it.t = t
	return it.t < timestamp.FromTime(it.end)
}

func (it *activityCountIterator) At() (int64, float64) {
	count := 0
	for _, activity := range it.activities {
		if time.Time(activity.StartDate).Before(timestamp.Time(it.t)) {
			count++
		} else {
			break
		}
	}
	return it.t, float64(count)
}

func (it *activityCountIterator) Next() bool {
	it.t += it.interval
	return it.t < timestamp.FromTime(it.end)
}

func (it *activityCountIterator) Err() error { return nil }
