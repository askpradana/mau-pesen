package entity

import "time"

type OTPCode struct {
	ID        uint      `gorm:"primaryKey"`
	Phone     string    `gorm:"type:text;not null"`
	Code      string    `gorm:"type:text;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	Used      bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
