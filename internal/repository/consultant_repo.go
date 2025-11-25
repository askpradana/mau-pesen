package repository

import (
	"nfldyprdn/maupesen/internal/entity"

	"gorm.io/gorm"
)

type ConsultantWithUser struct {
	entity.Consultant
	User entity.User `gorm:"foreignKey:UserID;references:ID"`
}

type ConsultantRepository interface {
	Create(c *entity.Consultant) error
	FindByUserID(userID string) (*entity.Consultant, error)
	Update(c *entity.Consultant) error
	ListActive() ([]ConsultantWithUser, error)
	ListAll() ([]ConsultantWithUser, error)
}

type consultantRepo struct {
	db *gorm.DB
}

func NewConsultantRepository(db *gorm.DB) ConsultantRepository {
	return &consultantRepo{db}
}

func (r *consultantRepo) Create(c *entity.Consultant) error {
	return r.db.Create(c).Error
}

func (r *consultantRepo) FindByUserID(userID string) (*entity.Consultant, error) {
	var c entity.Consultant
	err := r.db.Preload("User").First(&c, "user_id = ?", userID).Error
	return &c, err
}

func (r *consultantRepo) Update(c *entity.Consultant) error {
	return r.db.Save(c).Error
}

func (r *consultantRepo) ListActive() ([]ConsultantWithUser, error) {
	var list []ConsultantWithUser
	err := r.db.Joins("User").
		Where("consultants.is_active = true").
		Find(&list).Error
	return list, err
}

func (r *consultantRepo) ListAll() ([]ConsultantWithUser, error) {
	var list []ConsultantWithUser
	err := r.db.Joins("User").Find(&list).Error
	return list, err
}
