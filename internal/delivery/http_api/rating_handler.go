package http_api

import (
	"encoding/json"
	"net/http"
	"nfldyprdn/maupesen/internal/common"
	"nfldyprdn/maupesen/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func CreateRating(uc *usecase.RatingUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientID := r.Context().Value("user_id").(string)
		bookingID := chi.URLParam(r, "booking_id")
		var req struct {
			Rating  int    `json:"rating"`
			Message string `json:"message,omitempty"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			common.BadRequest(w, "Invalid JSON")
			return
		}
		if req.Rating < 1 || req.Rating > 5 {
			common.BadRequest(w, "Rating harus 1-5")
			return
		}
		if err := uc.Create(bookingID, clientID, req.Rating, req.Message); err != nil {
			common.BadRequest(w, err.Error())
			return
		}
		common.Created(w, "Terima kasih atas ratingnya! Rating berhasil disimpan.", nil)
	}
}
