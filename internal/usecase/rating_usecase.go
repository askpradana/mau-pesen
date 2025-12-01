package usecase

import (
	"fmt"
	"nfldyprdn/maupesen/internal/entity"
	"nfldyprdn/maupesen/internal/repository"
)

type RatingUsecase struct {
	ratingRepo  repository.RatingRepository
	bookingRepo repository.BookingRepository
}

func NewRatingUsecase(repos *repository.Repositories) *RatingUsecase {
	return &RatingUsecase{
		ratingRepo:  repos.Rating,
		bookingRepo: repos.Booking,
	}
}

func (u *RatingUsecase) Create(bookingID, clientID string, rating int, message string) error {
	booking, err := u.bookingRepo.FindByBookID(bookingID)

	if err != nil || booking == nil {
		return fmt.Errorf("booking ID not found")
	}

	if booking.Status != "approved" && booking.Status != "completed" {
		return fmt.Errorf("booking not completed/ cannot rating")
	}

	if booking.ClientID != clientID {
		return fmt.Errorf("it is not your booking")
	}

	exist, _ := u.ratingRepo.ExistsByBookingID(bookingID)

	if exist {
		return fmt.Errorf("rating already given")
	}

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
