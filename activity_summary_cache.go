package main

import (
	"encoding/gob"
	"os"
	"sync"
	"time"

	"github.com/csmarchbanks/stravaql/strava/model"
)

type activitySummaryCache struct {
	activities []*model.SummaryActivity
	diskCache  string
	cacheEnd   time.Time
	mux        sync.RWMutex
}

func newActivitySummaryCache(cacheFile string) *activitySummaryCache {
	activities := []*model.SummaryActivity{}
	cacheEnd := time.Unix(0, 0)
	if cacheFile != "" {
		file, err := os.Open(cacheFile)
		if err == nil {
			decoder := gob.NewDecoder(file)
			err = decoder.Decode(&activities)
			if err == nil {
				_ = decoder.Decode(&cacheEnd)
			}
		}
	}
	return &activitySummaryCache{
		activities: activities,
		diskCache:  cacheFile,
		cacheEnd:   cacheEnd,
	}
}

func (c *activitySummaryCache) Put(activities []*model.SummaryActivity) error {
	// Only cache results that are older than 7 days ago.
	newestToCache := time.Now().Add(-7 * 24 * time.Hour)

	toCache := []*model.SummaryActivity{}
	for _, activity := range activities {
		if time.Time(activity.StartDate).After(newestToCache) {
			continue
		}
		toCache = append(toCache, activity)
	}

	c.mux.RLock()
	if len(c.activities) == len(toCache) {
		c.mux.RUnlock()
		return nil
	}
	c.mux.RUnlock()
	c.mux.Lock()
	defer c.mux.Unlock()
	c.activities = toCache
	c.cacheEnd = newestToCache

	if c.diskCache != "" {
		file, err := os.Create(c.diskCache)
		if err != nil {
			return err
		}
		encoder := gob.NewEncoder(file)
		err = encoder.Encode(c.activities)
		if err != nil {
			return err
		}
		return encoder.Encode(c.cacheEnd)
	}

	return nil
}

func (c *activitySummaryCache) Get() ([]*model.SummaryActivity, time.Time) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.activities, c.cacheEnd
}
