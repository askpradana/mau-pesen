package entity

import "time"

type User struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:text;not null"`
	Email     string    `gorm:"type:text;uniqueIndex;not null"`
	Phone     string    `gorm:"type:text;uniqueIndex;not null"`
	Role      string    `gorm:"type:text;default:'client';index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
