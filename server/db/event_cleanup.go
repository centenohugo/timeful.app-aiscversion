package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"schej.it/server/models"
)

// How long past an event's last possible meeting time the retention window starts
// counting. A day only event stores midnight of each day with no duration, so its
// last day is still going for a full day (longer, west of UTC) after the instant
// stored for it — this keeps that from expiring the event a day early.
const expiryGracePeriod = 24 * time.Hour

// DeleteExpiredEvents permanently deletes every event whose last possible meeting
// time is older than retention, along with its responses, attendees and folder
// associations. Events that aren't pinned to real calendar dates (day of week
// events, availability groups) are never deleted — they have no last date to
// expire. Returns the number of events deleted.
func DeleteExpiredEvents(now time.Time, retention time.Duration) (int, error) {
	ctx := context.Background()
	cutoff := now.Add(-retention)

	// Narrow to events whose dates are all older than the cutoff. An event's real
	// expiry (last slot + duration + grace) is always later than its last date, so
	// this is a superset of the events that end up being deleted.
	candidates, err := EventsCollection.Find(ctx, bson.M{
		"type":    models.SPECIFIC_DATES,
		"dates.0": bson.M{"$exists": true},
		"dates":   bson.M{"$not": bson.M{"$gte": primitive.NewDateTimeFromTime(cutoff)}},
	}, options.Find().SetProjection(bson.M{
		"_id":          1,
		"type":         1,
		"dates":        1,
		"times":        1,
		"duration":     1,
		"signUpBlocks": 1,
	}))
	if err != nil {
		return 0, err
	}

	var events []models.Event
	if err := candidates.All(ctx, &events); err != nil {
		return 0, err
	}

	expiredIds := make([]primitive.ObjectID, 0)
	for _, event := range events {
		end, ok := event.LastPossibleEnd()
		if !ok {
			continue
		}
		if end.Add(expiryGracePeriod).Before(cutoff) {
			expiredIds = append(expiredIds, event.Id)
		}
	}
	if len(expiredIds) == 0 {
		return 0, nil
	}

	// Delete everything that hangs off the event first — an event left behind gets
	// picked up by the next run, whereas orphaned responses never would be
	idFilter := bson.M{"eventId": bson.M{"$in": expiredIds}}
	for _, collection := range []*mongo.Collection{
		EventResponsesCollection,
		AttendeesCollection,
		FolderEventsCollection,
	} {
		if _, err := collection.DeleteMany(ctx, idFilter); err != nil {
			return 0, err
		}
	}

	result, err := EventsCollection.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": expiredIds}})
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}
