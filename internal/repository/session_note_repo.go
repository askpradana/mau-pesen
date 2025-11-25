package repository

import (
	"nfldyprdn/maupesen/internal/entity"

	"gorm.io/gorm"
)

type SessionNoteRepository interface {
	Create(note *entity.SessionNote) error
	GetByBookingID(bookingID string) (*entity.SessionNote, error)
}

type sessionNoteRepo struct {
	db *gorm.DB
}

func NewSessionNoteRepository(db *gorm.DB) SessionNoteRepository {
	return &sessionNoteRepo{db}
}

func (r *sessionNoteRepo) Create(note *entity.SessionNote) error {
	return r.db.Create(note).Error
}

func (r *sessionNoteRepo) GetByBookingID(bookingID string) (*entity.SessionNote, error) {
	var note entity.SessionNote
	err := r.db.Preload("Consultant").First(&note, "booking_id = ?", bookingID).Error
	return &note, err
}
