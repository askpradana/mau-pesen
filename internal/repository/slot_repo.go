package repository

import (
	"nfldyprdn/maupesen/internal/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SlotRepository interface {
	CreateBatch(slots []*entity.ConsultantSlot) error
	GetAvailable(consultantID, date string) ([]entity.ConsultantSlot, error)
	LockAndBook(tx *gorm.DB, slotID string) (*entity.ConsultantSlot, error)
}

type slotRepo struct {
	db *gorm.DB
}

func NewSlotRepository(db *gorm.DB) SlotRepository {
	return &slotRepo{db}
}

func (r *slotRepo) CreateBatch(slots []*entity.ConsultantSlot) error {
	return r.db.Create(slots).Error
}

func (r *slotRepo) GetAvailable(consultantID, date string) ([]entity.ConsultantSlot, error) {
	var slots []entity.ConsultantSlot
	err := r.db.Where("consultant_id = ? AND date = ? AND available = true", consultantID, date).
		Order("hour ASC").Find(&slots).Error
	return slots, err
}

func (r *slotRepo) LockAndBook(tx *gorm.DB, slotID string) (*entity.ConsultantSlot, error) {
	var slot entity.ConsultantSlot
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ? AND available = true").
		First(&slot).Error
	if err != nil {
		return nil, err
	}
	slot.Available = false
	return &slot, tx.Save(&slot).Error
}
