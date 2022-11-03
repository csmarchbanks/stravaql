package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/csmarchbanks/stravaql/strava"
	"github.com/go-kit/log/level"
	"github.com/go-openapi/runtime/client"
	"github.com/gorilla/mux"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/grafana/regexp"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/route"
	"github.com/prometheus/common/version"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/rules"
	"github.com/prometheus/prometheus/scrape"
	"github.com/prometheus/prometheus/storage"
	promv1 "github.com/prometheus/prometheus/web/api/v1"
	v1 "github.com/prometheus/prometheus/web/api/v1"
)

func main() {
	promlogConfig := &promlog.Config{}
	flag.AddFlags(kingpin.CommandLine, promlogConfig)
	prometheus.MustRegister(version.NewCollector("stravaql"))
	kingpin.Version(version.Print("stravaql"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	// Initialize observability constructs.
	logger := promlog.New(promlogConfig)

	c := strava.NewHTTPClient(nil)
	token, err := oauthToken().Token()
	if err != nil {
		fmt.Println("unable to get oauth token", err)
		os.Exit(1)
	}
	auth := client.BearerToken(token.AccessToken)

	activityCache := newActivitySummaryCache(path.Join(".stravaql", "activity_cache.gob"))

	engineOpts := promql.EngineOpts{
		Logger: logger,
		Reg:    prometheus.DefaultRegisterer,
		// TODO: Make these configurable.
		MaxSamples:         10000000,
		Timeout:            2 * time.Minute,
		ActiveQueryTracker: nil,
		LookbackDelta:      5 * time.Minute,
		NoStepSubqueryIntervalFn: func(int64) int64 {
			return (15 * time.Second).Milliseconds()
		},
		EnableAtModifier:     true,
		EnableNegativeOffset: true,
	}

	promapi := promv1.NewAPI(
		promql.NewEngine(engineOpts),
		&sampleAndChunkQueryable{newQueryable(logger, c, auth, activityCache)},
		nil,
		nil,
		func(context.Context) promv1.TargetRetriever { return &emptyTargetRetriever{} },
		func(context.Context) promv1.AlertmanagerRetriever { return &emptyAlertmanagerRetriever{} },
		func() config.Config { return config.Config{} },
		map[string]string{}, // TODO: configuration flags?
		promv1.GlobalURLOptions{},
		func(f http.HandlerFunc) http.HandlerFunc { return f },
		nil,   // Only needed for admin APIs.
		"",    // This is for snapshots, which is disabled when admin APIs are disabled. Hence empty.
		false, // Disable admin APIs.
		logger,
		func(context.Context) v1.RulesRetriever { return &emptyRulesRetriever{} },
		0, 0, 0, // Remote read samples and concurrency limit.
		false, // isAgent flag set to false.
		regexp.MustCompile(".*"),
		func() (v1.RuntimeInfo, error) { return v1.RuntimeInfo{}, errors.New("not implemented") },
		&v1.PrometheusVersion{},
		// This is used for the stats API which we should not support. Or find other ways to.
		prometheus.GathererFunc(func() ([]*dto.MetricFamily, error) { return nil, nil }),
		prometheus.DefaultRegisterer,
		nil, // Will use defaultStatsRenderer if this is nil.
	)

	router := mux.NewRouter()
	promRouter := route.New().WithPrefix("/api/v1")
	promapi.Register(promRouter)

	// Register each path individually in order to get metrics about each endpoint.
	router.Path("/api/v1/metadata").Handler(promRouter)
	router.Path("/api/v1/read").Handler(promRouter)
	router.Path("/api/v1/read").Methods("POST").Handler(promRouter)
	router.Path("/api/v1/query").Methods("GET", "POST").Handler(promRouter)
	router.Path("/api/v1/query_range").Methods("GET", "POST").Handler(promRouter)
	router.Path("/api/v1/labels").Methods("GET", "POST").Handler(promRouter)
	router.Path("/api/v1/label/{name}/values").Methods("GET").Handler(promRouter)
	router.Path("/api/v1/series").Methods("GET", "POST", "DELETE").Handler(promRouter)
	router.Path("/api/v1/metadata").Methods("GET").Handler(promRouter)

	level.Info(logger).Log("msg", "serving StravaQL on :8055")
	if err := http.ListenAndServe(":8055", router); err != http.ErrServerClosed {
		level.Error(logger).Log("msg", "error serving metrics", "err", err)
	}
}

type sampleAndChunkQueryable struct {
	storage.Queryable
}

func (q *sampleAndChunkQueryable) ChunkQuerier(ctx context.Context, mint, maxt int64) (storage.ChunkQuerier, error) {
	return nil, errors.New("ChunkQuerier not implemented")
}

type emptyTargetRetriever struct{}

func (emptyTargetRetriever) TargetsActive() map[string][]*scrape.Target {
	return map[string][]*scrape.Target{}
}
func (emptyTargetRetriever) TargetsDropped() map[string][]*scrape.Target {
	return map[string][]*scrape.Target{}
}

type emptyAlertmanagerRetriever struct{}

func (emptyAlertmanagerRetriever) Alertmanagers() []*url.URL        { return nil }
func (emptyAlertmanagerRetriever) DroppedAlertmanagers() []*url.URL { return nil }

type emptyRulesRetriever struct{}

func (emptyRulesRetriever) RuleGroups() []*rules.Group           { return nil }
func (emptyRulesRetriever) AlertingRules() []*rules.AlertingRule { return nil }
