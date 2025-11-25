package usecase

import (
	"crypto/rand"
	"fmt"
	"nfldyprdn/maupesen/internal/config"
	"nfldyprdn/maupesen/internal/entity"
	"nfldyprdn/maupesen/internal/external/redis"
	"nfldyprdn/maupesen/internal/repository"
	"nfldyprdn/maupesen/pkg/jwt"
	"time"
)

type AuthUsecase struct {
	userRepo repository.UserRepository
	otpRepo  repository.OTPRepository
	logRepo  repository.LogRepository
}

func NewAuthUsecase(repos *repository.Repositories) *AuthUsecase {
	return &AuthUsecase{
		userRepo: repos.User,
		otpRepo:  repos.OTP,
		logRepo:  repos.Log,
	}
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Role         string `json:"role"`
}

// Register Client Biasa
func (u *AuthUsecase) RegisterClient(name, phone string) (*entity.User, error) {
	// Cek apakah phone sudah terdaftar
	if _, err := u.userRepo.FindByPhone(phone); err == nil {
		return nil, fmt.Errorf("nomor sudah terdaftar")
	}

	user := &entity.User{
		Name:  name,
		Phone: phone,
		Role:  "client",
	}

	if err := u.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Log registrasi
	u.logRepo.Create(nil, "user_registered", map[string]interface{}{
		"user_id": user.ID,
		"phone":   phone,
		"role":    "client",
	})

	return user, nil
}

// Register Admin (hanya via secret)
func (u *AuthUsecase) RegisterAdmin(name, phone, secret string) error {
	if secret != config.C.Admin.RegisterSecret {
		return fmt.Errorf("secret salah")
	}

	user := &entity.User{
		Name:  name,
		Phone: phone,
		Role:  "admin",
	}
	return u.userRepo.Create(user)
}

// Kirim OTP (dev mode: return OTP, production: kirim SMS)
func (u *AuthUsecase) RequestOTP(phone string) (string, error) {
	// Generate 6 digit OTP
	otp := fmt.Sprintf("%06d", randInt(100000, 999999))

	otpRecord := &entity.OTPCode{
		Phone:     phone,
		Code:      otp,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	if err := u.otpRepo.Create(otpRecord); err != nil {
		return "", err
	}

	// Di production: kirim SMS via Twilio/Nexmo
	// Di dev: return OTP langsung
	return otp, nil
}

// Verifikasi OTP + Login
func (u *AuthUsecase) VerifyOTP(phone, code string) (*LoginResponse, error) {
	otp, err := u.otpRepo.FindValid(phone, code)
	if err != nil {
		return nil, fmt.Errorf("OTP salah atau kadaluarsa")
	}

	// Tandai OTP sudah dipakai
	u.otpRepo.MarkUsed(otp.ID)

	user, err := u.userRepo.FindByPhone(phone)
	if err != nil {
		// Auto register kalau belum ada (optional)
		user, _ = u.RegisterClient("User "+phone[6:], phone)
	}

	if user.Role == "banned" {
		return nil, fmt.Errorf("akun anda diblokir")
	}

	accessToken, _ := jwt.Generate(user.ID, user.Role)
	refreshToken := generateRefreshToken()

	// Simpan refresh token di Redis (30 hari)
	redis.Client.Set(redis.Ctx, "refresh:"+refreshToken, user.ID, 30*24*time.Hour)

	// Log login
	u.logRepo.Create(&user.ID, "user_login", map[string]interface{}{
		"method": "otp",
		"phone":  phone,
	})

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user.ID,
		Name:         user.Name,
		Phone:        user.Phone,
		Role:         user.Role,
	}, nil
}

// Refresh Token
func (u *AuthUsecase) RefreshToken(oldRefreshToken string) (*LoginResponse, error) {
	userID, err := redis.Client.Get(redis.Ctx, "refresh:"+oldRefreshToken).Result()
	if err != nil {
		return nil, fmt.Errorf("refresh token invalid")
	}

	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	newAccess, _ := jwt.Generate(user.ID, user.Role)
	newRefresh := generateRefreshToken()

	redis.Client.Set(redis.Ctx, "refresh:"+newRefresh, user.ID, 30*24*time.Hour)
	redis.Client.Del(redis.Ctx, "refresh:"+oldRefreshToken)

	return &LoginResponse{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
		UserID:       user.ID,
		Name:         user.Name,
		Phone:        user.Phone,
		Role:         user.Role,
	}, nil
}

// Logout
func (u *AuthUsecase) Logout(refreshToken string) error {
	return redis.Client.Del(redis.Ctx, "refresh:"+refreshToken).Err()
}

// Helper
func randInt(min, max int) int {
	b := make([]byte, 8)
	rand.Read(b)
	return min + int(b[0])%(max-min+1)
}

func generateRefreshToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
