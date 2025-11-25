package usecase

import (
	"nfldyprdn/maupesen/internal/entity"
	"nfldyprdn/maupesen/internal/repository"
)

type ConsultantUsecase struct {
	consultantRepo repository.ConsultantRepository
}

func NewConsultantUsecase(repos *repository.Repositories) *ConsultantUsecase {
	return &ConsultantUsecase{
		consultantRepo: repos.Consultant,
	}
}

func (u *ConsultantUsecase) ListActive() ([]repository.ConsultantWithUser, error) {
	return u.consultantRepo.ListActive()
}

func (u *ConsultantUsecase) GetProfile(userID string) (*entity.Consultant, error) {
	return u.consultantRepo.FindByUserID(userID)
}
