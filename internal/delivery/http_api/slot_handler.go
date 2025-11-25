package http_api

import (
	"encoding/json"
	"net/http"
	"nfldyprdn/maupesen/internal/common"
	"nfldyprdn/maupesen/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func CreateBatchSlots(uc *usecase.SlotUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(string)
		var req struct {
			Date       string `json:"date"` // YYYY-MM-DD
			StartHours int    `json:"start_hours"`
			EndHours   int    `json:"end_hours"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			common.BadRequest(w, "invalid request")
			return
		}

		if req.StartHours < 0 || req.EndHours > 23 || req.StartHours >= req.EndHours {
			common.BadRequest(w, "invalid time")
		}

		if err := uc.CreateBatch(userID, req.Date, req.StartHours, req.EndHours); err != nil {
			common.BadRequest(w, err.Error())
			return
		}
		common.Created(w, "Slot berhasil dibuat", nil)
	}
}

func GetAvailableSlots(uc *usecase.SlotUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		consultantID := chi.URLParam(r, "id")
		date := r.URL.Query().Get("date")
		slots, err := uc.GetAvailable(consultantID, date)
		if err != nil {
			common.BadRequest(w, "Slot tidak ditemukan")
			return
		}
		common.OK(w, "Slot tersedia", slots)
	}
}
