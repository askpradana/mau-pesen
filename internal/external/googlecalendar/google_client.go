package googlecalendar

import (
	"context"
	"fmt"
	"nfldyprdn/maupesen/internal/config"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var Service *calendar.Service

func NewClient() {
	ctx := context.Background()
	srv, _ := calendar.NewService(ctx, option.WithCredentialsFile(config.C.GoogleCalendar.ServiceAccountFile))
	Service = srv
}

func CreateEvent(consultantEmail, clientName, purpose, date string, hour int) (string, error) {
	start := fmt.Sprintf("%sT%02d:00:00+07:00", date, hour)
	end := fmt.Sprintf("%sT%02d:00:00+07:00", date, hour+1)

	event := &calendar.Event{
		Summary:     fmt.Sprintf("Konsultasi: %s", purpose),
		Description: purpose,
		Start:       &calendar.EventDateTime{DateTime: start, TimeZone: "Asia/Jakarta"},
		End:         &calendar.EventDateTime{DateTime: end, TimeZone: "Asia/Jakarta"},
		Attendees:   []*calendar.EventAttendee{{Email: consultantEmail}},
		ConferenceData: &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestId:             fmt.Sprintf("maupesen-%d", time.Now().UnixNano()),
				ConferenceSolutionKey: &calendar.ConferenceSolutionKey{Type: "hangoutsMeet"},
			},
		},
	}

	event, err := Service.Events.Insert(config.C.GoogleCalendar.CalendarID, event).
		ConferenceDataVersion(1).SendUpdates("all").Do()
	if err != nil {
		return "", err
	}
	return event.Id, nil
}
