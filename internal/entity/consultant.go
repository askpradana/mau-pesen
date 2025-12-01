package entity

import "time"

type Consultant struct {
	UserID      string    `gorm:"primaryKey;type:uuid"`
	Name        string    `gorm:"type:text"`
	Phone       string    `gorm:"type:text;uniqueIndex;not null"`
	Email       string    `gorm:"type:text;uniqueIndex;not null"`
	Speciality  string    `gorm:"type:text"`
	Bio         string    `gorm:"type:text"`
	RatingAvg   float64   `gorm:"type:numeric(3,2);default:0"`
	TotalRating int       `gorm:"default:0"`
	Price       int       `gorm:"type:integer;column:price_per_session;default:0"`
	IsActive    bool      `gorm:"default:true"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`

	User User `gorm:"foreignKey:UserID;references:ID"`
}
