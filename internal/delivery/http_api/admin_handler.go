package http_api

import (
	"encoding/json"
	"net/http"
	"nfldyprdn/maupesen/internal/common"
	"nfldyprdn/maupesen/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func AdminCreateConsultant(uc *usecase.AdminUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, reqW *http.Request) {
		var req struct {
			UserID     string `json:"user_id"`
			Speciality string `json:"speciality"`
			Bio        string `json:"bio"`
			Price      int    `json:"price"`
		}
		if err := json.NewDecoder(reqW.Body).Decode(&req); err != nil {
			common.BadRequest(w, "Invalid JSON")
			return
		}
		if err := uc.CreateConsultant(req.UserID, req.Speciality, req.Bio, req.Price); err != nil {
			common.BadRequest(w, err.Error())
			return
		}
		common.Created(w, "Consultant berhasil dibuat", nil)
	}
}

// PUT /admin/consultants/{id} → update
func AdminUpdateConsultant(uc *usecase.AdminUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "id")
		var req struct {
			Speciality string `json:"speciality"`
			Bio        string `json:"bio"`
			Price      int    `json:"price"`
			IsActive   bool   `json:"is_active"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		err := uc.UpdateConsultant(userID, req.Speciality, req.Bio, req.Price, req.IsActive)
		if err != nil {
			common.BadRequest(w, err.Error())
			return
		}
		common.OK(w, "Consultant updated", nil)
	}
}

// DELETE /admin/consultants/{id} → deactivate
func AdminDeactivateConsultant(uc *usecase.AdminUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "id")
		if err := uc.DeactivateConsultant(userID); err != nil {
			common.BadRequest(w, err.Error())
			return
		}
		common.OK(w, "Consultant dinonaktifkan", nil)
	}
}

// GET /admin/consultants → List semua consultant (aktif + nonaktif)
func AdminListConsultants(uc *usecase.AdminUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		consultants, err := uc.ListAllConsultants()
		if err != nil {
			common.InternalError(w, "Gagal mengambil data consultant")
			return
		}
		common.OK(w, "Daftar semua consultant", consultants)
	}
}

// GET /admin/bookings → Lihat semua booking di sistem
func AdminListAllBookings(uc *usecase.AdminUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookings, err := uc.ListAllBookings()
		if err != nil {
			common.InternalError(w, "Gagal mengambil data booking")
			return
		}
		common.OK(w, "Semua booking di sistem", bookings)
	}
}

// POST /admin/ban/{user_id} → Ban user nakal (client atau consultant)
func AdminBanUser(uc *usecase.AdminUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "user_id")
		if userID == "" {
			common.BadRequest(w, "user_id wajib diisi")
			return
		}

		var req struct {
			Reason string `json:"reason,omitempty"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			common.BadRequest(w, "Invalid JSON")
			return
		}

		// Ambil admin yang lagi login
		adminID, ok := r.Context().Value("user_id").(string)
		if !ok || adminID == "" {
			common.Unauthorized(w, "admin tidak terdeteksi")
			return
		}

		// BAN USER + LOG SEKALIGUS (pake method yang bener)
		if err := uc.BanUser(userID, req.Reason, adminID); err != nil {
			common.BadRequest(w, err.Error())
			return
		}

		common.OK(w, "User berhasil dibanned", map[string]any{
			"user_id":    userID,
			"status":     "banned",
			"banned_by":  adminID,
			"ip_address": r.RemoteAddr,
			"user_agent": r.UserAgent(),
		})
	}
}
