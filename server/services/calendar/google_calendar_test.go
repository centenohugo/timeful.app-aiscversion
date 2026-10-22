package calendar

import (
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestBuildGoogleCalendarDraftURL(t *testing.T) {
	start := time.Date(2026, 10, 23, 15, 0, 0, 0, time.FixedZone("CEST", 2*60*60))
	end := start.Add(90 * time.Minute)

	draftUrl := BuildGoogleCalendarDraftURL(
		"Team sync & planning",
		"Scheduled with Meet AISC",
		start,
		end,
		"Europe/Madrid",
		[]string{"a@example.com", "b@example.com"},
	)

	if !strings.HasPrefix(draftUrl, "https://calendar.google.com/calendar/render?") {
		t.Fatalf("unexpected base URL: %s", draftUrl)
	}
	parsed, err := url.Parse(draftUrl)
	if err != nil {
		t.Fatalf("draft URL does not parse: %v", err)
	}
	q := parsed.Query()

	want := map[string]string{
		"action":  "TEMPLATE",
		"text":    "Team sync & planning",
		"details": "Scheduled with Meet AISC",
		"dates":   "20261023T130000Z/20261023T143000Z",
		"ctz":     "Europe/Madrid",
		"add":     "a@example.com,b@example.com",
		"vcon":    "meet",
	}
	for key, value := range want {
		if got := q.Get(key); got != value {
			t.Errorf("%s = %q, want %q", key, got, value)
		}
	}
}

func TestBuildGoogleCalendarDraftURLWithoutGuests(t *testing.T) {
	start := time.Date(2026, 10, 23, 13, 0, 0, 0, time.UTC)
	draftUrl := BuildGoogleCalendarDraftURL("Sync", "", start, start.Add(time.Hour), "", nil)

	parsed, err := url.Parse(draftUrl)
	if err != nil {
		t.Fatalf("draft URL does not parse: %v", err)
	}
	q := parsed.Query()
	if q.Has("add") {
		t.Errorf("expected no add param, got %q", q.Get("add"))
	}
	if q.Has("ctz") {
		t.Errorf("expected no ctz param, got %q", q.Get("ctz"))
	}
	if q.Get("vcon") != "meet" {
		t.Errorf("vcon = %q, want meet", q.Get("vcon"))
	}
}
