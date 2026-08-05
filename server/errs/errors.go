package errs

import (
	"encoding/json"
	"fmt"
)

// Errors enum
// TODO: make these an actual type (i.e. Errors.NotSignedIn)
const (
	NotSignedIn           string = "not-signed-in"
	UserDoesNotExist      string = "user-does-not-exist"
	EventNotFound         string = "event-not-found"
	FriendRequestNotFound string = "friend-request-not-found"
	UserNotFriends        string = "user-not-friends"
	UserNotEventOwner     string = "user-not-event-owner"
	AttendeeEmailNotFound string = "attendee-email-not-found"
	EventNotGroup         string = "event-not-group"
	InvalidCredentials    string = "invalid-credentials"
	InvalidIdToken        string = "invalid-id-token"
	InvalidProxyName      string = "invalid-proxy-name"
	ProxyNotAllowed       string = "proxy-not-allowed"
	InvalidScheduler      string = "invalid-scheduler"

	NoGoogleCalendarAccount         string = "no-google-calendar-account"
	CalendarWritePermissionRequired string = "calendar-write-permission-required"
	FailedToCreateCalendarEvent     string = "failed-to-create-calendar-event"
)

type GoogleAPIError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Details interface{} `json:"details"`
	Errors  interface{} `json:"errors"`
}

func (e *GoogleAPIError) Error() string {
	s, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintln("GoogleAPIError: <error parsing json>")
	}

	return fmt.Sprintln("GoogleAPIError: ", string(s))
}
