package usecase

import (
	"fmt"
	"nfldyprdn/maupesen/internal/entity"
	"nfldyprdn/maupesen/internal/repository"
	"time"

	"gorm.io/gorm"
)

type BookingUsecase struct {
	bookingRepo repository.BookingRepository
	slotRepo    repository.SlotRepository
	logRepo     repository.LogRepository
	db          *gorm.DB
}

func NewBookingUsecase(repos *repository.Repositories) *BookingUsecase {
	return &BookingUsecase{
		bookingRepo: repos.Booking,
		slotRepo:    repos.Slot,
		logRepo:     repos.Log,
		db:          repos.DB,
	}
}

func (u *BookingUsecase) Create(clientID, slotID, purpose string) (string, error) {
	var bookingID string
	err := u.db.Transaction(func(tx *gorm.DB) error {
		slot, err := u.slotRepo.LockAndBook(tx, slotID)
		if err != nil {
			return fmt.Errorf("slot sudah dibooking orang lain")
		}

		bookingID = fmt.Sprintf("BOOK-%s", time.Now().Format("20060102150405"))
		booking := &entity.Booking{
			BookingID:    bookingID,
			ClientID:     clientID,
			ConsultantID: slot.ConsultantID,
			SlotID:       slot.ID,
			Date:         slot.Date,
			Hour:         slot.Hour,
			Purpose:      purpose,
			Status:       entity.Pending,
		}
		return u.bookingRepo.Create(tx, booking)
	})
	if err == nil {
		u.logRepo.Create(&clientID, "booking_created", map[string]any{"booking_id": bookingID})
	}
	return bookingID, err
}

func (u *BookingUsecase) Approve(bookingID string) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		b, err := u.bookingRepo.FindByBookID(bookingID)
		if err != nil {
			return fmt.Errorf("booking not found")
		}

		if b.Status != entity.Pending {
			return fmt.Errorf("booking invalid")
		}
		//_, _, err := googlecalendar.CreateEvent(*b.ConsultantID, "Konsultasi", b.Purpose, b.Date, b.Hour)
		if err != nil {
			return err
		}
		b.Status = entity.Approved
		//b.GoogleEventID = &eventID
		// Simpan meet link di metadata atau kirim notif (nanti)
		return u.bookingRepo.Update(tx, b)
	})
}

func (u *BookingUsecase) Reject(bookingID, reason string) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		b, err := u.bookingRepo.FindByBookID(bookingID)
		if err != nil {
			return fmt.Errorf("booking not found")
		}

		if b.Status != entity.Pending {
			return fmt.Errorf("booking invalid")
		}
		b.Status = entity.Rejected
		b.RejectReason = &reason
		return u.bookingRepo.Update(tx, b)
	})
}
