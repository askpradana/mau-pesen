package repository

import (
	"nfldyprdn/maupesen/internal/entity"
	"time"

	"gorm.io/gorm"
)

type OTPRepository interface {
	Create(otp *entity.OTPCode) error
	FindValid(phone, code string) (*entity.OTPCode, error)
	MarkUsed(id uint) error
}

type otpRepo struct {
	db *gorm.DB
}

func NewOTPRepository(db *gorm.DB) OTPRepository {
	return &otpRepo{db}
}

func (r *otpRepo) Create(otp *entity.OTPCode) error {
	return r.db.Create(otp).Error
}

func (r *otpRepo) FindValid(phone, code string) (*entity.OTPCode, error) {
	var otp entity.OTPCode
	err := r.db.Where("phone = ? AND code = ? AND expires_at > ? AND used = false",
		phone, code, time.Now()).First(&otp).Error
	return &otp, err
}

func (r *otpRepo) MarkUsed(id uint) error {
	return r.db.Model(&entity.OTPCode{}).Where("id = ?", id).Update("used", true).Error
}
