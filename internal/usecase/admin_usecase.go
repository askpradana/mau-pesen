package usecase

import (
	"fmt"
	"nfldyprdn/maupesen/internal/entity"
	"nfldyprdn/maupesen/internal/repository"
)

type AdminUsecase struct {
	userRepo       repository.UserRepository
	consultantRepo repository.ConsultantRepository
	bookingRepo    repository.BookingRepository
	logRepo        repository.LogRepository
	authUseCase    *AuthUsecase
}

func NewAdminUsecase(repos *repository.Repositories, auth *AuthUsecase) *AdminUsecase {
	return &AdminUsecase{
		userRepo:       repos.User,
		consultantRepo: repos.Consultant,
		bookingRepo:    repos.Booking,
		logRepo:        repos.Log,
		authUseCase:    auth,
	}
}

// 1. Create Consultant
func (a *AdminUsecase) CreateConsultant(name, phone, speciality, bio string, price int) error {

	auth, err := a.authUseCase.RegisterClient(name, phone)
	//user, err := a.userRepo.FindByID(auth.Phone)
	//if err != nil || user.Role != "client" {
	//	return fmt.Errorf("user tidak ditemukan atau sudah punya role")
	//}

	if err != nil {
		return fmt.Errorf("user already exists")
	}

	auth.Role = "consultant"
	a.userRepo.Update(auth)

	consultant := &entity.Consultant{
		UserID:     auth.ID,
		Name:       name,
		Phone:      phone,
		Speciality: speciality,
		Bio:        bio,
		Price:      price,
		IsActive:   true,
	}
	return a.consultantRepo.Create(consultant)
}

// 2. Update Consultant
func (a *AdminUsecase) UpdateConsultant(userID, speciality, bio string, price int, isActive bool) error {
	c, err := a.consultantRepo.FindByUserID(userID)
	if err != nil {
		return err
	}
	c.Speciality = speciality
	c.Bio = bio
	c.Price = price
	c.IsActive = isActive
	return a.consultantRepo.Update(c)
}

// 3. Delete / Nonaktifkan Consultant
func (a *AdminUsecase) DeactivateConsultant(userID string) error {
	c, err := a.consultantRepo.FindByUserID(userID)
	if err != nil {
		return err
	}
	c.IsActive = false
	return a.consultantRepo.Update(c)
}

// 4. List Semua Consultant (aktif & nonaktif)
func (a *AdminUsecase) ListAllConsultants() ([]repository.ConsultantWithUser, error) {
	return a.consultantRepo.ListAll()
}

// 5. Ban User (client/consultant)
func (a *AdminUsecase) BanUser(userID, reason, bannedBy string) error {
	user, err := a.userRepo.FindByID(userID)
	if err != nil {
		return err
	}
	user.Role = "banned"
	if err := a.userRepo.Update(user); err != nil {
		return err
	}

	a.logRepo.Create(&bannedBy, "user_banned", map[string]interface{}{
		"target_user_id": userID,
		"reason":         reason,
		"banned_by":      bannedBy,
	})

	return nil
}

// 6. View All Bookings (admin bisa lihat semua)
func (a *AdminUsecase) ListAllBookings() ([]entity.Booking, error) {
	return a.bookingRepo.ListAll()
}
