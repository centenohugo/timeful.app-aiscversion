package calendar

import (
	"time"

	"schej.it/server/models"
)

type CalendarProvider interface {
	GetCalendarList() (map[string]models.SubCalendar, error)
	GetCalendarEvents(calendarId string, timeMin time.Time, timeMax time.Time) ([]models.CalendarEvent, error)
}

func GetCalendarProvider(calendarAccount models.CalendarAccount) CalendarProvider {
	if calendarAccount.CalendarType == models.GoogleCalendarType {
		return &GoogleCalendar{
			OAuth2CalendarAuth: *calendarAccount.OAuth2CalendarAuth,
		}
	}
	return nil
}
