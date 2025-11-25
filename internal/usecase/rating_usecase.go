package usecase

import (
	"nfldyprdn/maupesen/internal/entity"
	"nfldyprdn/maupesen/internal/repository"
)

type RatingUsecase struct {
	ratingRepo repository.RatingRepository
}

func NewRatingUsecase(repos *repository.Repositories) *RatingUsecase {
	return &RatingUsecase{
		ratingRepo: repos.Rating,
	}
}

func (u *RatingUsecase) Create(bookingID, clientID string, rating int, message string) error {
	r := &entity.Rating{
		BookingID: bookingID,
		ClientID:  clientID,
		Rating:    rating,
		Message:   message,
	}
	if err := u.ratingRepo.Create(r); err != nil {
		return err
	}
	// Auto update consultant rating_avg
	return u.ratingRepo.UpdateConsultantAvg(bookingID)
}
