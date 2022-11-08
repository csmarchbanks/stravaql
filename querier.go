package main

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/csmarchbanks/stravaql/strava"
	"github.com/csmarchbanks/stravaql/strava/activities"
	"github.com/csmarchbanks/stravaql/strava/model"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-openapi/runtime/client"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/timestamp"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb/chunkenc"
	"github.com/prometheus/prometheus/tsdb/tsdbutil"
	"golang.org/x/oauth2"
)

const (
	activitiesMetric      = "strava_activities_total"
	movingDurationMetric  = "strava_activity_moving_duration_seconds_total"
	elapsedDurationMetric = "strava_activity_elapsed_duration_seconds_total"
	distanceMetric        = "strava_activity_distance_meters_total"
	elevationGainMetric   = "strava_activity_elevation_gain_meters_total"
)

var metrics = []string{
	activitiesMetric,
	movingDurationMetric,
	elapsedDurationMetric,
	distanceMetric,
	elevationGainMetric,
}

type querier struct {
	logger        log.Logger
	ctx           context.Context
	client        *strava.StravaAPIV3
	tokenSource   oauth2.TokenSource
	activityCache *activitySummaryCache
	mint, maxt    int64
}

func newQueryable(logger log.Logger, sc *strava.StravaAPIV3, tokenSource oauth2.TokenSource, activityCache *activitySummaryCache) storage.Queryable {
	return storage.QueryableFunc(func(ctx context.Context, mint, maxt int64) (storage.Querier, error) {
		return &querier{
			logger:        logger,
			ctx:           ctx,
			client:        sc,
			tokenSource:   tokenSource,
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

	var warnings storage.Warnings
	allActivities, fetchStart := q.activityCache.Get()

	token, err := q.tokenSource.Token()
	if err != nil {
		level.Error(q.logger).Log("msg", "unable to get oauth token", "err", err)
		warnings = append(warnings, fmt.Errorf("unable to get oauth token: %w", err))
	} else {
		pageSize := int64(100)
		page := int64(1)
		after := fetchStart.Unix()
		params := activities.NewGetLoggedInAthleteActivitiesParamsWithContext(ctx).
			WithDefaults().
			WithPerPage(&pageSize).
			WithPage(&page).
			WithAfter(&after)

		for {
			activities, err := q.client.Activities.GetLoggedInAthleteActivities(params, client.BearerToken(token.AccessToken))
			if err != nil {
				warnings = append(warnings, err)
				break
			}
			if len(activities.Payload) == 0 {
				break
			}
			allActivities = append(allActivities, activities.Payload...)
			page++
		}
	}
	sort.Slice(allActivities, func(i, j int) bool {
		a := time.Time(allActivities[i].StartDate)
		b := time.Time(allActivities[j].StartDate)
		return a.Before(b)
	})

	if len(warnings) > 0 {
		var sb strings.Builder
		for i, warning := range warnings {
			sb.WriteString(warning.Error())
			if i < len(warnings)-1 {
				sb.WriteString(", ")
			}
		}
		level.Warn(q.logger).Log("msg", "warnings while querying", "warnings", sb.String())
	}

	if err := q.activityCache.Put(allActivities); err != nil {
		level.Warn(q.logger).Log("msg", "error putting activities in cache", "err", err)
	}

	switch metricName {
	case activitiesMetric:
		return newActivitySummationSeriesSet(allActivities, metricName, activityCount, q.mint, q.maxt, interval, matchers, warnings)
	case movingDurationMetric:
		return newActivitySummationSeriesSet(allActivities, metricName, activityMovingDuration, q.mint, q.maxt, interval, matchers, warnings)
	case elapsedDurationMetric:
		return newActivitySummationSeriesSet(allActivities, metricName, activityElapsedDuration, q.mint, q.maxt, interval, matchers, warnings)
	case distanceMetric:
		return newActivitySummationSeriesSet(allActivities, metricName, activityDistance, q.mint, q.maxt, interval, matchers, warnings)
	case elevationGainMetric:
		return newActivitySummationSeriesSet(allActivities, metricName, activityElevationGain, q.mint, q.maxt, interval, matchers, warnings)
	default:
		return storage.EmptySeriesSet()
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

func activityDistance(activity *model.SummaryActivity) float64 {
	return float64(activity.Distance)
}

func activityElevationGain(activity *model.SummaryActivity) float64 {
	return float64(activity.TotalElevationGain)
}

func (q *querier) LabelValues(name string, matchers ...*labels.Matcher) ([]string, storage.Warnings, error) {
	switch name {
	case "__name__":
		return metrics, nil, nil
	case "sport":
		return []string{
			string(model.SportTypeAlpineSki),
			string(model.SportTypeBackcountrySki),
			string(model.SportTypeCanoeing),
			string(model.SportTypeCrossfit),
			string(model.SportTypeEBikeRide),
			string(model.SportTypeElliptical),
			string(model.SportTypeEMountainBikeRide),
			string(model.SportTypeGolf),
			string(model.SportTypeGravelRide),
			string(model.SportTypeHandcycle),
			string(model.SportTypeHike),
			string(model.SportTypeIceSkate),
			string(model.SportTypeInlineSkate),
			string(model.SportTypeKayaking),
			string(model.SportTypeKitesurf),
			string(model.SportTypeMountainBikeRide),
			string(model.SportTypeNordicSki),
			string(model.SportTypeRide),
			string(model.SportTypeRockClimbing),
			string(model.SportTypeRollerSki),
			string(model.SportTypeRowing),
			string(model.SportTypeRun),
			string(model.SportTypeSail),
			string(model.SportTypeSkateboard),
			string(model.SportTypeSnowboard),
			string(model.SportTypeSnowshoe),
			string(model.SportTypeSoccer),
			string(model.SportTypeStairStepper),
			string(model.SportTypeStandUpPaddling),
			string(model.SportTypeSurfing),
			string(model.SportTypeSwim),
			string(model.SportTypeTrailRun),
			string(model.SportTypeVelomobile),
			string(model.SportTypeVirtualRide),
			string(model.SportTypeVirtualRun),
			string(model.SportTypeWalk),
			string(model.SportTypeWeightTraining),
			string(model.SportTypeWheelchair),
			string(model.SportTypeWindsurf),
			string(model.SportTypeWorkout),
			string(model.SportTypeYoga),
		}, nil, nil
	default:
		return nil, nil, nil
	}
}

func (q *querier) LabelNames(matchers ...*labels.Matcher) ([]string, storage.Warnings, error) {
	return []string{"__name__", "sport"}, nil, nil
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
	warnings         storage.Warnings

	i int
}

func newActivitySummationSeriesSet(activities []*model.SummaryActivity, metricName string, summationFn activitySummationFn, start, end, interval int64, matchers []*labels.Matcher, warnings storage.Warnings) storage.SeriesSet {
	activitiesByType := map[string][]*model.SummaryActivity{}
	sportTypes := []string{}

activityLoop:
	for _, activity := range activities {
		sportType := string(activity.SportType)
		if sportType == "" {
			sportType = string(activity.Type)
		}

		if activitiesByType[sportType] == nil {
			// Check if this sport type is not included by any matchers.
			for _, matcher := range matchers {
				if matcher.Name == "sport" && !matcher.Matches(sportType) {
					continue activityLoop
				}
			}

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
		warnings:         warnings,
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
	return ss.warnings
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
	start := s.start.UnixMilli()
	end := s.end.UnixMilli()
	samples := make(tsdbutil.SampleSlice, 0, (end-start)/s.interval+1)
	value := 0.0
	i := 0

	for ts := start; ts <= end; ts += s.interval {
		for ; i < len(s.activities); i++ {
			activity := s.activities[i]
			if time.Time(activity.StartDate).UnixMilli() <= ts {
				value += s.summationFn(activity)
			} else {
				break
			}
		}

		samples = append(samples, newSample(ts, value))
	}

	return storage.NewListSeriesIterator(samples)
}

type sample struct {
	t int64
	v float64
}

func newSample(t int64, v float64) tsdbutil.Sample {
	return sample{t, v}
}

func (s sample) T() int64 {
	return s.t
}

func (s sample) V() float64 {
	return s.v
}
