package middleware_api

import (
	"context"
	"net/http"
	"nfldyprdn/maupesen/internal/common"
	"nfldyprdn/maupesen/pkg/jwt"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			common.Unauthorized(w, "Token required")
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		claims, err := jwt.Validate(token)
		if err != nil {
			common.Unauthorized(w, "Invalid token")
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "role", claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
