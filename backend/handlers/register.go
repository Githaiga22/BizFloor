package handlers

import (
	"net/http"

	"gorm.io/gorm"
)

func RegisterHandlers(mux *http.ServeMux, db *gorm.DB) {
	profileHandler := &ProfileHandler{DB: db}
	serviceHandler := &ServiceHandler{DB: db}
	bookingHandler := &BookingHandler{DB: db}
	paymentHandler := &PaymentHandler{DB: db}

	mux.HandleFunc("/profile", profileHandler.CreateProfile)
	mux.HandleFunc("/add-service", serviceHandler.AddServiceHandler)
	mux.HandleFunc("/create-booking", bookingHandler.CreateBookingHandler)
	mux.HandleFunc("/create-booking/pay", paymentHandler.CreatePaymentHandler)
}
