package repository

import (
	"nfldyprdn/maupesen/internal/entity"

	"gorm.io/gorm"
)

type BookingRepository interface {
	Create(tx *gorm.DB, b *entity.Booking) error
	FindByBookID(id string) (*entity.Booking, error)
	Update(tx *gorm.DB, b *entity.Booking) error
	ListAll() ([]entity.Booking, error)
}

type bookingRepo struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepo{db}
}

func (r *bookingRepo) Create(tx *gorm.DB, b *entity.Booking) error {
	return tx.Create(b).Error
}

func (r *bookingRepo) FindByBookID(bookId string) (*entity.Booking, error) {
	var b entity.Booking
	err := r.db.Preload("Client").Preload("Consultant").Preload("Slot").First(&b, "booking_id = ?", bookId).Error
	return &b, err
}

func (r *bookingRepo) Update(tx *gorm.DB, b *entity.Booking) error {
	return tx.Save(b).Error
}

func (r *bookingRepo) ListAll() ([]entity.Booking, error) {
	var list []entity.Booking
	err := r.db.Preload("Client").Preload("Consultant").Find(&list).Error
	return list, err
}
