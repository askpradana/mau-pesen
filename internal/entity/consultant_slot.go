package entity

import "time"

type ConsultantSlot struct {
	ID           string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	SlotID       *string   `gorm:"type:text;uniqueIndex"`
	ConsultantID string    `gorm:"type:uuid;not null"`
	Date         string    `gorm:"type:date;not null"`
	Hour         int       `gorm:"not null;check:hour >= 0 AND hour <= 23"`
	Available    bool      `gorm:"default:true"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	User User `gorm:"foreignKey:UserID;references:ID"`
}
