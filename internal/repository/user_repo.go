package repository

import (
	"nfldyprdn/maupesen/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByID(id string) (*entity.User, error)
	FindByPhone(phone string) (*entity.User, error)
	Update(user *entity.User) error
	Ban(userID string) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindByID(id string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByPhone(phone string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "phone = ?", phone).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepo) Ban(userID string) error {
	return r.db.Model(&entity.User{}).
		Where("id = ?", userID).
		Update("role", "banned").Error
}
