package handlers

import (
	"gorm.io/gorm"
	"encoding/json"
	"net/http"
)

type Booking struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	ClientID uint   `json:"client_id"`
	Service  string `json:"service"`
	Date     string `json:"date"`
	Status   string `json:"status"`
}

func GetBookingsHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := r.URL.Query().Get("status")
		var bookings []Booking

		// Query bookings based on status filter
		if status != "" {
			db.Where("status = ?", status).Find(&bookings)
		} else {
			db.Find(&bookings)
		}

		// Encode response as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bookings)
	}
}
