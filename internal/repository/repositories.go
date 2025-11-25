package repository

import "gorm.io/gorm"

type Repositories struct {
	User        UserRepository
	Consultant  ConsultantRepository
	Slot        SlotRepository
	Booking     BookingRepository
	Rating      RatingRepository
	SessionNote SessionNoteRepository
	Log         LogRepository
	OTP         OTPRepository
	DB          *gorm.DB
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:        NewUserRepository(db),
		Consultant:  NewConsultantRepository(db),
		Slot:        NewSlotRepository(db),
		Booking:     NewBookingRepository(db),
		Rating:      NewRatingRepository(db),
		SessionNote: NewSessionNoteRepository(db),
		Log:         NewLogRepository(db),
		OTP:         NewOTPRepository(db),
		DB:          db,
	}
}
