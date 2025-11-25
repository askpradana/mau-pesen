package entity

import "time"

type SessionNote struct {
	BookingID    string    `gorm:"type:uuid;primaryKey"`
	ConsultantID string    `gorm:"type:uuid"`
	Notes        string    `gorm:"type:text"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	// relasi (optional, biar Preload jalan)
	Booking    Booking `gorm:"foreignKey:BookingID;references:ID"`
	Consultant User    `gorm:"foreignKey:ConsultantID;references:ID"`
}
