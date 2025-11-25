package usecase

type Usecases struct {
	Auth       *AuthUsecase
	Consultant *ConsultantUsecase
	Slot       *SlotUsecase
	Booking    *BookingUsecase
	Rating     *RatingUsecase
	Admin      *AdminUsecase
}
