package entity

import "time"

type Rating struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BookingID string    `gorm:"type:text;uniqueIndex;not null"`
	ClientID  string    `gorm:"type:uuid;not null"`
	Rating    int       `gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Message   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	// Relasi
	Booking Booking `gorm:"foreignKey:BookingID;references:BookingID"`
	Client  User    `gorm:"foreignKey:ClientID;references:ID"`
}
