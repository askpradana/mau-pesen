package http_api

import (
	"encoding/json"
	"net/http"
	"nfldyprdn/maupesen/internal/common"
	"nfldyprdn/maupesen/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func CreateBooking(uc *usecase.BookingUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientID := r.Context().Value("user_id").(string)
		var req struct {
			SlotID  string `json:"slot_id"`
			Purpose string `json:"purpose"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			common.BadRequest(w, "Invalid JSON")
			return
		}
		bookingID, err := uc.Create(clientID, req.SlotID, req.Purpose)
		if err != nil {
			common.BadRequest(w, err.Error())
			return
		}
		common.Created(w, "Booking berhasil dibuat", map[string]string{"booking_id": bookingID})
	}
}

func ApproveBooking(uc *usecase.BookingUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookingID := chi.URLParam(r, "id")
		if err := uc.Approve(bookingID); err != nil {
			common.BadRequest(w, err.Error())
			return
		}
		common.OK(w, "Booking disetujui & link Google Meet terkirim", nil)
	}
}

func RejectBooking(uc *usecase.BookingUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookingID := chi.URLParam(r, "id")
		var req struct {
			Reason string `json:"reason,omitempty"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		if err := uc.Reject(bookingID, req.Reason); err != nil {
			common.BadRequest(w, err.Error())
			return
		}
		common.OK(w, "booking ditolak", nil)
	}
}
