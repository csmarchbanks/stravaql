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

	allActivities, fetchStart := q.activityCache.Get()
	pageSize := int64(100)
	page := int64(1)
	after := fetchStart.Unix()
	params := activities.NewGetLoggedInAthleteActivitiesParamsWithContext(ctx).
		WithDefaults().
		WithPerPage(&pageSize).
		WithPage(&page).
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

	switch metricName {
	case "strava_activities_total":
		return newActivitySummationSeriesSet(allActivities, "strava_activities_total", activityCount, q.mint, q.maxt, interval)
	case "strava_activity_moving_duration_seconds_total":
		return newActivitySummationSeriesSet(allActivities, "strava_activity_moving_duration_seconds_total", activityMovingDuration, q.mint, q.maxt, interval)
	case "strava_activity_elapsed_duration_seconds_total":
		return newActivitySummationSeriesSet(allActivities, "strava_activity_elapsed_duration_seconds_total", activityElapsedDuration, q.mint, q.maxt, interval)
	default:
		return nil
	}
}

func activityCount(activity *model.SummaryActivity) float64 {
	return 1.0
}

func activityMovingDuration(activity *model.SummaryActivity) float64 {
	return float64(activity.MovingTime)
}

func activityElapsedDuration(activity *model.SummaryActivity) float64 {
	return float64(activity.ElapsedTime)
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

type activitySummationFn func(*model.SummaryActivity) float64

type activitySummationSeriesSet struct {
	activitiesByType map[string][]*model.SummaryActivity
	sportTypes       []string
	metricName       string
	summationFn      activitySummationFn
	start, end       time.Time
	interval         int64

	i int
}

func newActivitySummationSeriesSet(activities []*model.SummaryActivity, metricName string, summationFn activitySummationFn, start, end, interval int64) storage.SeriesSet {
	activitiesByType := map[string][]*model.SummaryActivity{}
	sportTypes := []string{}
	for _, activity := range activities {
		sportType := string(activity.SportType)
		if sportType == "" {
			sportType = string(activity.Type)
		}
		if activitiesByType[sportType] == nil {
			sportTypes = append(sportTypes, sportType)
		}
		activitiesByType[sportType] = append(activitiesByType[sportType], activity)
	}

	return &activitySummationSeriesSet{
		activitiesByType: activitiesByType,
		sportTypes:       sportTypes,
		metricName:       metricName,
		summationFn:      summationFn,
		start:            timestamp.Time(start),
		end:              timestamp.Time(end),
		interval:         interval,
		i:                -1,
	}
}

func (ss *activitySummationSeriesSet) Next() bool {
	ss.i++
	return ss.i < len(ss.sportTypes)
}

func (ss *activitySummationSeriesSet) At() storage.Series {
	sportType := ss.sportTypes[ss.i]
	return &activitySummationSeries{
		activities: ss.activitiesByType[sportType],
		labels: labels.FromStrings(
			"__name__", ss.metricName,
			"sport", sportType,
		),
		summationFn: ss.summationFn,
		start:       ss.start,
		end:         ss.end,
		interval:    ss.interval,
	}
}

func (ss *activitySummationSeriesSet) Err() error {
	return nil
}

func (ss *activitySummationSeriesSet) Warnings() storage.Warnings {
	return nil
}

type activitySummationSeries struct {
	activities  []*model.SummaryActivity
	labels      labels.Labels
	summationFn activitySummationFn
	start, end  time.Time
	interval    int64
}

func (s *activitySummationSeries) Labels() labels.Labels {
	return s.labels
}

func (s *activitySummationSeries) Iterator() chunkenc.Iterator {
	return &activityCountIterator{
		activities:  s.activities,
		summationFn: s.summationFn,
		start:       s.start,
		end:         s.end,
		interval:    s.interval,
		t:           timestamp.FromTime(s.start),
	}
}

type activityCountIterator struct {
	activities  []*model.SummaryActivity
	summationFn activitySummationFn
	start, end  time.Time
	interval    int64
	t           int64
}

func (it *activityCountIterator) Seek(t int64) bool {
	it.t = t
	return it.t < timestamp.FromTime(it.end)
}

func (it *activityCountIterator) At() (int64, float64) {
	value := 0.0
	for _, activity := range it.activities {
		if time.Time(activity.StartDate).Before(timestamp.Time(it.t)) {
			value += it.summationFn(activity)
		} else {
			break
		}
	}
	return it.t, value
}

func (it *activityCountIterator) Next() bool {
	it.t += it.interval
	return it.t < timestamp.FromTime(it.end)
}

func (it *activityCountIterator) Err() error { return nil }
