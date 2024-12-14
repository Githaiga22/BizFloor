package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"group7/models"
	"gorm.io/gorm"
)

type BookingHandler struct {
	DB *gorm.DB
}

func NewBookingHandler(db *gorm.DB) *BookingHandler {
	return &BookingHandler{
		DB: db,
	}
}

type BookingRequest struct {
	ServiceID   uint      `json:"service_id"`
	BookingTime time.Time `json:"booking_time"`
}

func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get user ID from context (set by auth middleware)
	userID := r.Context().Value("user_id").(uint)

	// Get service details and validate
	var service models.Service
	if err := h.DB.Preload("Business").First(&service, req.ServiceID).Error; err != nil {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}

	// Check for double booking
	var existingBooking models.Booking
	if err := h.DB.Where("service_id = ? AND booking_time = ? AND status != ?",
		req.ServiceID, req.BookingTime, "cancelled").First(&existingBooking).Error; err == nil {
		http.Error(w, "Time slot already booked", http.StatusConflict)
		return
	}

	// Create booking
	booking := models.Booking{
		ServiceID:   req.ServiceID,
		ClientID:    userID,
		BusinessID:  service.BusinessID,
		BookingTime: req.BookingTime,
		Status:      "pending",
	}

	// Start a transaction
	tx := h.DB.Begin()

	if err := tx.Create(&booking).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to create booking", http.StatusInternalServerError)
		return
	}

	// Create initial payment record
	payment := models.Payment{
		BookingID:   booking.ID,
		Amount:      service.Price,
		Status:      "pending",
		PaymentTime: time.Now(),
	}

	if err := tx.Create(&payment).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to create payment record", http.StatusInternalServerError)
		return
	}

	// Update booking with payment ID
	booking.PaymentID = &payment.ID
	if err := tx.Save(&booking).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to update booking with payment", http.StatusInternalServerError)
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	// Load relationships for response
	h.DB.Preload("Service").Preload("Client").Preload("Business").Preload("Payment").First(&booking, booking.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

func (h *BookingHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.Context().Value("user_id").(uint)

	// Get user to check if they're a business owner
	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	query := h.DB.Model(&models.Booking{}).
		Preload("Service").
		Preload("Client").
		Preload("Business").
		Preload("Payment")

	if user.IsBusinessOwner {
		// For business owners, show bookings for their business
		query = query.Where("business_id IN (SELECT id FROM businesses WHERE owner_id = ?)", userID)
	} else {
		// For regular users, show their bookings
		query = query.Where("client_id = ?", userID)
	}

	var bookings []models.Booking
	if err := query.Find(&bookings).Error; err != nil {
		http.Error(w, "Failed to fetch bookings", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

func (h *BookingHandler) UpdateBookingStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updateReq struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	bookingID := r.URL.Query().Get("id")
	if bookingID == "" {
		http.Error(w, "Booking ID required", http.StatusBadRequest)
		return
	}

	var booking models.Booking
	if err := h.DB.First(&booking, bookingID).Error; err != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	// Update booking status
	booking.Status = updateReq.Status
	if err := h.DB.Save(&booking).Error; err != nil {
		http.Error(w, "Failed to update booking", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

func (h *BookingHandler) GetAvailableSlots(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse query parameters
	date := r.URL.Query().Get("date")
	serviceID := r.URL.Query().Get("service_id")

	if date == "" || serviceID == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	// Parse the date
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	// Get service details to know duration
	var service models.Service
	if err := h.DB.First(&service, serviceID).Error; err != nil {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}

	// Generate time slots for the day (assuming 9 AM to 5 PM)
	startTime := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 9, 0, 0, 0, time.Local)
	endTime := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 17, 0, 0, 0, time.Local)

	// Get existing bookings for this service on this day
	var existingBookings []models.Booking
	h.DB.Where("service_id = ? AND DATE(booking_time) = ? AND status != ?", 
		serviceID, date, "cancelled").Find(&existingBookings)

	// Create map of booked times for quick lookup
	bookedTimes := make(map[time.Time]bool)
	for _, booking := range existingBookings {
		bookedTimes[booking.BookingTime] = true
	}

	// Generate available slots
	var slots []struct {
		Time      time.Time `json:"time"`
		Available bool      `json:"available"`
	}

	for t := startTime; t.Before(endTime); t = t.Add(time.Duration(service.DurationMins) * time.Minute) {
		slots = append(slots, struct {
			Time      time.Time `json:"time"`
			Available bool      `json:"available"`
		}{
			Time:      t,
			Available: !bookedTimes[t],
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slots)
}