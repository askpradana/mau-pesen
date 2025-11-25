package entity

import (
	"encoding/json"
	"time"
)

type Log struct {
	ID        uint            `gorm:"primaryKey"`
	UserID    *string         `gorm:"type:uuid"`
	Activity  string          `gorm:"type:text;not null"`
	Metadata  json.RawMessage `gorm:"type:jsonb"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`

	// Relasi (optional)
	User *User `gorm:"foreignKey:UserID;references:ID"`
}
