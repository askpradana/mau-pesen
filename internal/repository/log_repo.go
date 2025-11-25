package repository

import (
	"encoding/json"
	"nfldyprdn/maupesen/internal/entity"

	"gorm.io/gorm"
)

type LogRepository interface {
	Create(userID *string, activity string, metadata map[string]any) error
}

type logRepo struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepo{db}
}

func (r *logRepo) Create(userID *string, activity string, metadata map[string]any) error {
	metaBytes, _ := json.Marshal(metadata)
	log := &entity.Log{
		UserID:   userID,
		Activity: activity,
		Metadata: metaBytes,
	}
	return r.db.Create(log).Error
}
