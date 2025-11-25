package middleware_api

import (
	"net/http"
	"nfldyprdn/maupesen/internal/common"
)

func RequireRole(required string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role := r.Context().Value("role").(string)
			if role != required {
				common.Forbidden(w, "Akses ditolak")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
