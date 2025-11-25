package main

import (
	"fmt"
	"log"
	"net/http"
	"nfldyprdn/maupesen/internal/config"
	"nfldyprdn/maupesen/internal/database"
	"nfldyprdn/maupesen/internal/delivery/http_api"
	"nfldyprdn/maupesen/internal/external/googlecalendar"
	"nfldyprdn/maupesen/internal/external/redis"
	"nfldyprdn/maupesen/internal/middleware_api"
	"nfldyprdn/maupesen/internal/repository"
	"nfldyprdn/maupesen/internal/usecase"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	config.Init()
	database.Connect()
	redis.InitClient()
	googlecalendar.NewClient()

	repos := repository.NewRepositories(database.DB)
	uc := &usecase.Usecases{
		Auth:       usecase.NewAuthUsecase(repos),
		Consultant: usecase.NewConsultantUsecase(repos),
		Slot:       usecase.NewSlotUsecase(repos),
		Booking:    usecase.NewBookingUsecase(repos),
		Rating:     usecase.NewRatingUsecase(repos),
		Admin:      usecase.NewAdminUsecase(repos),
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("mau-pesen-ultimate API v2025 — PRODUCTION READY"))
	})

	// Public
	r.Post("/register", http_api.RegisterClient(uc.Auth))
	r.Post("/admin/register", http_api.RegisterAdmin(uc.Auth))
	r.Post("/otp/request", http_api.RequestOTP(uc.Auth))
	r.Post("/otp/verify", http_api.VerifyOTP(uc.Auth))
	r.Post("/token/refresh", http_api.RefreshToken(uc.Auth))
	r.Post("/logout", http_api.Logout(uc.Auth))

	// Protected
	r.Group(func(r chi.Router) {
		r.Use(middleware_api.AuthMiddleware)

		r.Get("/consultants", http_api.ListConsultants(uc.Consultant))
		r.Get("/consultants/me", http_api.GetConsultantProfile(uc.Consultant))
		r.Get("/consultants/{id}/slots", http_api.GetAvailableSlots(uc.Slot))
		r.Post("/slots/batch", http_api.CreateBatchSlots(uc.Slot))
		r.Post("/bookings", http_api.CreateBooking(uc.Booking))
		r.Post("/ratings/{booking_id}", http_api.CreateRating(uc.Rating))

		// Consultant only
		r.Group(func(r chi.Router) {
			r.Use(middleware_api.RequireRole("consultant"))
			r.Patch("/bookings/{id}/approve", http_api.ApproveBooking(uc.Booking))
			r.Patch("/bookings/{id}/reject", http_api.RejectBooking(uc.Booking))
		})

		// Admin only
		r.Group(func(r chi.Router) {
			r.Use(middleware_api.RequireRole("admin"))
			r.Post("/admin/consultants", http_api.AdminCreateConsultant(uc.Admin))
			r.Put("/admin/consultants/{id}", http_api.AdminUpdateConsultant(uc.Admin))
			r.Delete("/admin/consultants/{id}", http_api.AdminDeactivateConsultant(uc.Admin))
			r.Post("/admin/ban/{user_id}", http_api.AdminBanUser(uc.Admin))
			r.Get("/admin/consultants", http_api.AdminListConsultants(uc.Admin))
			r.Get("/admin/bookings", http_api.AdminListAllBookings(uc.Admin))
		})
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	//http_api.HandleFunc("/health", healthCheck)
	//
	//port := ":8080"
	//fmt.Printf("Server starting on port %s\n", port)
	//log.Fatal(http_api.ListenAndServe(port, nil))

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "ok", "message": "Online Booking System"}`)
}
