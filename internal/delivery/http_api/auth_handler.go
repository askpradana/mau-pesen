package http_api

import (
	"encoding/json"
	"net/http"
	"nfldyprdn/maupesen/internal/common"
	"nfldyprdn/maupesen/internal/usecase"
)

func RegisterClient(uc *usecase.AuthUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name  string `json:"name"`
			Phone string `json:"phone"`
			Email string `json:"email"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		user, err := uc.RegisterClient(req.Name, req.Phone, req.Email)
		if err != nil {
			common.BadRequest(w, err.Error())
			return
		}
		common.Created(w, "Registrasi client berhasil", user)
	}
}

func RegisterAdmin(uc *usecase.AuthUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name   string `json:"name"`
			Phone  string `json:"phone"`
			Email  string `json:"email"`
			Secret string `json:"secret"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		if err := uc.RegisterAdmin(req.Name, req.Phone, req.Email, req.Secret); err != nil {
			common.Forbidden(w, "error create admin")
			return
		}
		common.Created(w, "Admin berhasil dibuat", nil)
	}
}

func RequestOTP(uc *usecase.AuthUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Phone string `json:"phone"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		code, err := uc.RequestOTP(req.Phone)
		if err != nil {
			common.BadRequest(w, err.Error())
			return
		}
		common.OK(w, "OTP terkirim (dev mode)", map[string]string{"code": code})
	}
}

func VerifyOTP(uc *usecase.AuthUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Phone string `json:"phone"`
			Code  string `json:"code"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		resp, err := uc.VerifyOTP(req.Phone, req.Code)
		if err != nil {
			common.BadRequest(w, err.Error())
			return
		}
		common.OK(w, "Login berhasil", resp)
	}
}

func RefreshToken(uc *usecase.AuthUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			RefreshToken string `json:"refresh_token"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			common.BadRequest(w, "invalid json")
			return
		}
		resp, err := uc.RefreshToken(req.RefreshToken)
		if err != nil {
			common.Unauthorized(w, err.Error())
			return
		}
		common.OK(w, "token refreshed", resp)
	}
}

func Logout(uc *usecase.AuthUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			RefreshToken string `json:"refresh_token"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			common.BadRequest(w, "invalid json")
			return
		}
		if err := uc.Logout(req.RefreshToken); err != nil {
			common.BadRequest(w, err.Error())
			return
		}
		common.OK(w, "logout berhasil", nil)
	}
}
