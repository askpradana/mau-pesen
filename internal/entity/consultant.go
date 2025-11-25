package entity

import "time"

type Consultant struct {
	UserID      string  `gorm:"primaryKey;type:uuid"`
	Speciality  string  `gorm:"type:text"`
	Bio         string  `gorm:"type:text"`
	RatingAvg   float64 `gorm:"type:numeric(3,2);default:0"`
	TotalRating int     `gorm:"default:0"`
	Price       int
	IsActive    bool      `gorm:"default:true"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
