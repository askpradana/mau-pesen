package entity

import "time"

type BookingStatus string

const (
	Pending   BookingStatus = "pending"
	Approved  BookingStatus = "approved"
	Rejected  BookingStatus = "rejected"
	Cancelled BookingStatus = "cancelled"
)

type Booking struct {
	ID        string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BookingID string `gorm:"type:text;uniqueIndex;not null"`

	ClientID string `gorm:"type:uuid"`
	Client   User   `gorm:"foreignKey:ClientID;references:ID"`

	ConsultantID string `gorm:"type:uuid"`
	Consultant   User   `gorm:"foreignKey:ConsultantID;references:ID"`

	SlotID string         `gorm:"type:uuid"`
	Slot   ConsultantSlot `gorm:"foreignKey:SlotID;references:ID"`

	Date          string        `gorm:"type:date;not null"`
	Hour          int           `gorm:"not null"`
	Purpose       string        `gorm:"type:text"`
	Status        BookingStatus `gorm:"type:text;default:'pending'"`
	GoogleEventID *string
	RejectReason  *string
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
