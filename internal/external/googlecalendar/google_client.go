package googlecalendar

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"nfldyprdn/maupesen/internal/config"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var Service *calendar.Service
var once sync.Once

func Init() error {
	rand.Seed(time.Now().UnixNano())

	var err error
	once.Do(func() {
		err = initOAuth2()
	})
	if err != nil {
		return err
	}
	fmt.Println("GOOGLE CALENDAR SIAP — MEET LINK OTOMATIS AKTIF!")
	return nil
}

func initOAuth2() error {
	ctx := context.Background()

	// BACA CREDENTIALS.JSON (DARI OAUTH CLIENT ID - DESKTOP APP)
	credFile := config.C.GoogleCalendar.ServiceAccountFile
	b, err := os.ReadFile(credFile)
	if err != nil {
		return fmt.Errorf("credentials.json tidak ditemukan!\nDownload dari: https://console.cloud.google.com/apis/credentials → OAuth 2.0 Client ID → Desktop App → Download JSON → taruh di internal/credentials/")
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		return fmt.Errorf("gagal parse credentials.json: %v", err)
	}

	client := getOAuth2Client(config)

	var svcErr error
	Service, svcErr = calendar.NewService(ctx, option.WithHTTPClient(client))
	if svcErr != nil {
		return fmt.Errorf("gagal buat calendar service: %v", svcErr)
	}

	return nil
}

func getOAuth2Client(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil || !tok.Valid() {
		fmt.Println("Token tidak ada atau expired → minta izin akses...")
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
		fmt.Println("Token baru disimpan! Selanjutnya otomatis.")
	} else {
		fmt.Println("Token valid ditemukan → langsung pakai")
	}
	return config.Client(context.Background(), tok)
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	return t, err
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("maupesen-2025", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	fmt.Printf("\nBUKA LINK INI DI BROWSER (PAKE AKUN susanathena4@gmail.com):\n%s\n\nSetelah klik 'Allow', COPY KODE yang muncul → PASTE DI SINI ↓\nKODE: ", authURL)

	var code string
	if _, err := fmt.Scanln(&code); err != nil {
		panic("Gagal baca kode: " + err.Error())
	}

	tok, err := config.Exchange(context.Background(), code)
	if err != nil {
		panic("Gagal tukar kode jadi token: " + err.Error())
	}
	return tok
}

func saveToken(file string, token *oauth2.Token) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic("Gagal simpan token: " + err.Error())
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func CreateEvent(consultantEmail, consultantName, clientName, clientEmail, purpose, date string, hour int) (string, string, error) {
	if Service == nil {
		return "", "", fmt.Errorf("google calendar client never init, call googlecalendar.init() first")
	}

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return "", "", err
	}

	pureDate := strings.Split(date, "T")[0]
	y, m, d := parseYMD(pureDate)

	startT := time.Date(y, m, d, hour, 0, 0, 0, loc)
	endT := startT.Add(time.Hour)

	start := startT.Format(time.RFC3339)
	end := endT.Format(time.RFC3339)

	event := &calendar.Event{

		Summary: fmt.Sprintf(
			"Konsultasi Online - %s",
			consultantName),

		Description: fmt.Sprintf(
			"Pasien: %s\nKeluhan: %s\n\nDetail Konsultasi Online:\n• Harap bergabung 5–10 menit sebelum jadwal dimulai\n• Pastikan koneksi internet stabil\n• Gunakan perangkat dengan kamera & mikrofon\n\nJika Anda mengalami kendala atau ingin menjadwalkan ulang,\nsilakan hubungi Admin Telemed kami.",
			clientName,
			purpose),

		Start: &calendar.EventDateTime{
			DateTime: start,
			TimeZone: "Asia/Jakarta",
		},

		End: &calendar.EventDateTime{
			DateTime: end,
			TimeZone: "Asia/Jakarta",
		},

		Attendees: []*calendar.EventAttendee{
			{Email: consultantEmail, DisplayName: consultantName},
			{Email: clientEmail, DisplayName: clientName}},

		Reminders: &calendar.EventReminders{
			UseDefault: false,
			Overrides: []*calendar.EventReminder{
				{Method: "email", Minutes: 60},
				{Method: "popup", Minutes: 10},
			},
			ForceSendFields: []string{"UseDefault", "Overrides"},
		},

		ConferenceData: &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestId:             fmt.Sprintf("maupesen-%d-%d", time.Now().UnixNano(), rand.Intn(1000000)),
				ConferenceSolutionKey: &calendar.ConferenceSolutionKey{Type: "hangoutsMeet"},
			},
		},
	}

	createdEvent, err := Service.Events.Insert(config.C.GoogleCalendar.CalendarID, event).
		ConferenceDataVersion(1).
		SendUpdates("all").
		Do()

	if err != nil {
		return "", "", fmt.Errorf("failed create event: %v", err)
	}

	meetLink := createdEvent.HangoutLink
	if meetLink == "" {
		meetLink = "https://meet.google.com"
	}

	return createdEvent.Id, meetLink, nil
}

func parseYMD(date string) (int, time.Month, int) {
	parts := strings.Split(date, "-")
	y, _ := strconv.Atoi(parts[0])
	mInt, _ := strconv.Atoi(parts[1])
	d, _ := strconv.Atoi(parts[2])
	return y, time.Month(mInt), d
}
