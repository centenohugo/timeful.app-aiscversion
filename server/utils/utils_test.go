package utils

import "testing"

func TestIsValidEmail(t *testing.T) {
	valid := []string{
		"someone@example.com",
		"first.last+tag@sub.example.co.uk",
		"a@b.io",
	}
	for _, email := range valid {
		if !IsValidEmail(email) {
			t.Errorf("IsValidEmail(%q) = false, want true", email)
		}
	}

	// Guest respondents type this field themselves, so anything can end up in it. Google rejects
	// the whole event creation request if one attendee address is malformed.
	invalid := []string{
		"",
		"   ",
		"Hugo",
		"hugo@",
		"@example.com",
		"hugo example@test.com",
		"Hugo <hugo@example.com>",
		"one@example.com, two@example.com",
		"a@b@c.com",
	}
	for _, email := range invalid {
		if IsValidEmail(email) {
			t.Errorf("IsValidEmail(%q) = true, want false", email)
		}
	}
}

func TestHasCalendarWriteScope(t *testing.T) {
	writeScope := "openid email https://www.googleapis.com/auth/calendar.events"
	if !HasCalendarWriteScope(writeScope) {
		t.Errorf("HasCalendarWriteScope(%q) = false, want true", writeScope)
	}

	// calendar.events.readonly must not satisfy the write requirement.
	readOnly := "openid email https://www.googleapis.com/auth/calendar.events.readonly"
	if HasCalendarWriteScope(readOnly) {
		t.Errorf("HasCalendarWriteScope(%q) = true, want false", readOnly)
	}

	if HasCalendarWriteScope("") {
		t.Error("HasCalendarWriteScope(\"\") = true, want false")
	}
}
