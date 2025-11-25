package http_api

import (
	"net/http"
	"nfldyprdn/maupesen/internal/common"
	"nfldyprdn/maupesen/internal/usecase"
)

func ListConsultants(uc *usecase.ConsultantUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := uc.ListActive()
		if err != nil {
			common.InternalError(w, err.Error())
			return
		}
		common.OK(w, "daftar konsultan aktif", list)
	}
}

func GetConsultantProfile(uc *usecase.ConsultantUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(string)
		profile, err := uc.GetProfile(userID)
		if err != nil {
			common.NotFound(w, "profile tidak ditemukan")
			return
		}
		common.OK(w, "profile konsultan", profile)
	}
}
