// Package cleanup runs the background jobs that keep the database from filling up
// with data nobody can use anymore.
package cleanup

import (
	"os"
	"strconv"
	"time"

	"schej.it/server/db"
	"schej.it/server/logger"
)

// How long after an event's last possible meeting time it gets deleted, unless
// EVENT_RETENTION_DAYS says otherwise
const DefaultRetentionDays = 30

// How long after its creation a day of week event or availability group gets
// deleted, unless UNDATED_EVENT_RETENTION_DAYS says otherwise. These events have no
// last meeting date, so they're aged from when they were created instead.
const DefaultUndatedRetentionDays = 90

// How often expired events are swept up. Events expire on a scale of days, so
// there's nothing to gain from checking more often than daily.
const sweepInterval = 24 * time.Hour

// StartExpiredEventSweeper deletes expired events now and once a day from here on,
// and returns a function that stops the sweeper. Deleting dated events is off when
// EVENT_RETENTION_DAYS is zero or negative, and deleting day of week events and
// availability groups is off when UNDATED_EVENT_RETENTION_DAYS is.
func StartExpiredEventSweeper() func() {
	retention, enabled := retentionFromEnv("EVENT_RETENTION_DAYS", DefaultRetentionDays)
	if !enabled {
		logger.StdOut.Println("Expired event deletion is disabled (EVENT_RETENTION_DAYS <= 0)")
	}
	undatedRetention, undatedEnabled := retentionFromEnv("UNDATED_EVENT_RETENTION_DAYS", DefaultUndatedRetentionDays)
	if !undatedEnabled {
		logger.StdOut.Println("Stale day of week / group event deletion is disabled (UNDATED_EVENT_RETENTION_DAYS <= 0)")
	}
	if !enabled && !undatedEnabled {
		return func() {}
	}

	sweepAll := func() {
		if enabled {
			sweep(retention)
		}
		if undatedEnabled {
			sweepUndated(undatedRetention)
		}
	}

	stop := make(chan struct{})
	go func() {
		ticker := time.NewTicker(sweepInterval)
		defer ticker.Stop()

		sweepAll()
		for {
			select {
			case <-ticker.C:
				sweepAll()
			case <-stop:
				return
			}
		}
	}()

	return func() { close(stop) }
}

func sweep(retention time.Duration) {
	deleted, err := db.DeleteExpiredEvents(time.Now(), retention)
	if err != nil {
		logger.StdErr.Printf("Failed to delete expired events: %v\n", err)
		return
	}
	if deleted > 0 {
		logger.StdOut.Printf("Deleted %d event(s) that ended more than %d days ago\n", deleted, int(retention.Hours()/24))
	}
}

func sweepUndated(retention time.Duration) {
	deleted, err := db.DeleteStaleUndatedEvents(time.Now(), retention)
	if err != nil {
		logger.StdErr.Printf("Failed to delete stale day of week / group events: %v\n", err)
		return
	}
	if deleted > 0 {
		logger.StdOut.Printf("Deleted %d day of week / group event(s) created more than %d days ago\n", deleted, int(retention.Hours()/24))
	}
}

// retentionFromEnv returns the retention window configured by the given env var
// (or defaultDays if it's unset), and whether that kind of deletion is enabled at all
func retentionFromEnv(envVar string, defaultDays int) (time.Duration, bool) {
	days := defaultDays

	if value := os.Getenv(envVar); len(value) > 0 {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			logger.StdErr.Printf("Invalid %s %q, defaulting to %d days\n", envVar, value, defaultDays)
		} else {
			days = parsed
		}
	}

	if days <= 0 {
		return 0, false
	}

	return time.Duration(days) * 24 * time.Hour, true
}
