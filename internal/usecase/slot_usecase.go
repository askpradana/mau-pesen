package usecase

import (
	"nfldyprdn/maupesen/internal/entity"
	"nfldyprdn/maupesen/internal/repository"
)

type SlotUsecase struct {
	slotRepo repository.SlotRepository
}

func NewSlotUsecase(repos *repository.Repositories) *SlotUsecase {
	return &SlotUsecase{
		slotRepo: repos.Slot,
	}
}

func (u *SlotUsecase) CreateBatch(consultantID, date string, startHour, endHour int) error {
	var slots []*entity.ConsultantSlot
	for h := startHour; h < endHour; h++ {
		slots = append(slots, &entity.ConsultantSlot{
			ConsultantID: consultantID,
			Date:         date,
			Hour:         h,
			Available:    true,
		})
	}
	return u.slotRepo.CreateBatch(slots)
}

func (u *SlotUsecase) GetAvailable(consultantID, date string) ([]entity.ConsultantSlot, error) {
	return u.slotRepo.GetAvailable(consultantID, date)
}
