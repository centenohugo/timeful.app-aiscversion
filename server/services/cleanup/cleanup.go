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

// How often expired events are swept up. Events expire on a scale of days, so
// there's nothing to gain from checking more often than daily.
const sweepInterval = 24 * time.Hour

// StartExpiredEventSweeper deletes expired events now and once a day from here on,
// and returns a function that stops the sweeper. Deletion is off entirely when
// EVENT_RETENTION_DAYS is zero or negative.
func StartExpiredEventSweeper() func() {
	retention, enabled := retentionFromEnv()
	if !enabled {
		logger.StdOut.Println("Expired event deletion is disabled (EVENT_RETENTION_DAYS <= 0)")
		return func() {}
	}

	stop := make(chan struct{})
	go func() {
		ticker := time.NewTicker(sweepInterval)
		defer ticker.Stop()

		sweep(retention)
		for {
			select {
			case <-ticker.C:
				sweep(retention)
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

// retentionFromEnv returns the configured retention window, and whether deleting
// expired events is enabled at all
func retentionFromEnv() (time.Duration, bool) {
	days := DefaultRetentionDays

	if value := os.Getenv("EVENT_RETENTION_DAYS"); len(value) > 0 {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			logger.StdErr.Printf("Invalid EVENT_RETENTION_DAYS %q, defaulting to %d days\n", value, DefaultRetentionDays)
		} else {
			days = parsed
		}
	}

	if days <= 0 {
		return 0, false
	}

	return time.Duration(days) * 24 * time.Hour, true
}
