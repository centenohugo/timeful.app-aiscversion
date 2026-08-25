package models_test

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"schej.it/server/models"
)

func date(t *testing.T, value string) primitive.DateTime {
	t.Helper()

	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		t.Fatalf("bad test date %q: %v", value, err)
	}

	return primitive.NewDateTimeFromTime(parsed)
}

func duration(hours float32) *float32 {
	return &hours
}

func TestLastPossibleEnd(t *testing.T) {
	blockEnd := date(t, "2024-03-12T23:00:00Z")

	tests := []struct {
		name     string
		event    models.Event
		expected string
		ok       bool
	}{
		{
			name: "last date plus duration",
			event: models.Event{
				Type:     models.SPECIFIC_DATES,
				Dates:    []primitive.DateTime{date(t, "2024-03-12T09:00:00Z"), date(t, "2024-03-10T09:00:00Z")},
				Duration: duration(8),
			},
			expected: "2024-03-12T17:00:00Z",
			ok:       true,
		},
		{
			name: "fractional duration",
			event: models.Event{
				Type:     models.SPECIFIC_DATES,
				Dates:    []primitive.DateTime{date(t, "2024-03-12T09:00:00Z")},
				Duration: duration(1.5),
			},
			expected: "2024-03-12T10:30:00Z",
			ok:       true,
		},
		{
			name: "days only event has no duration",
			event: models.Event{
				Type:     models.SPECIFIC_DATES,
				Dates:    []primitive.DateTime{date(t, "2024-03-12T00:00:00Z")},
				Duration: duration(0),
			},
			expected: "2024-03-12T00:00:00Z",
			ok:       true,
		},
		{
			name: "specific times can run past the last date",
			event: models.Event{
				Type:     models.SPECIFIC_DATES,
				Dates:    []primitive.DateTime{date(t, "2024-03-12T09:00:00Z")},
				Times:    []primitive.DateTime{date(t, "2024-03-12T20:00:00Z")},
				Duration: duration(1),
			},
			expected: "2024-03-12T21:00:00Z",
			ok:       true,
		},
		{
			name: "sign up blocks count towards the end",
			event: models.Event{
				Type:         models.SPECIFIC_DATES,
				Dates:        []primitive.DateTime{date(t, "2024-03-12T09:00:00Z")},
				Duration:     duration(1),
				SignUpBlocks: &[]models.SignUpBlock{{EndDate: &blockEnd}},
			},
			expected: "2024-03-13T00:00:00Z",
			ok:       true,
		},
		{
			name: "day of week events never end",
			event: models.Event{
				Type:     models.DOW,
				Dates:    []primitive.DateTime{date(t, "2018-06-17T09:00:00Z")},
				Duration: duration(8),
			},
			ok: false,
		},
		{
			name: "availability groups never end",
			event: models.Event{
				Type:     models.GROUP,
				Dates:    []primitive.DateTime{date(t, "2018-06-17T09:00:00Z")},
				Duration: duration(8),
			},
			ok: false,
		},
		{
			name: "event without dates",
			event: models.Event{
				Type:     models.SPECIFIC_DATES,
				Duration: duration(8),
			},
			ok: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			end, ok := test.event.LastPossibleEnd()
			if ok != test.ok {
				t.Fatalf("expected ok to be %v, got %v", test.ok, ok)
			}
			if !test.ok {
				return
			}

			expected := date(t, test.expected).Time()
			if !end.Equal(expected) {
				t.Errorf("expected %v, got %v", expected.UTC(), end.UTC())
			}
		})
	}
}
