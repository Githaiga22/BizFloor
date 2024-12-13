// File: handlers/service_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"group7/models"

	"gorm.io/gorm"
)

type ServiceHandler struct {
	DB *gorm.DB
}

func (h *ServiceHandler) AddServiceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var service models.Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the BusinessID
	var business models.Business
	if err := h.DB.First(&business, service.BusinessID).Error; err != nil {
		http.Error(w, "Business not found", http.StatusNotFound)
		return
	}

	// Set timestamps for creation
	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()

	if err := h.DB.Create(&service).Error; err != nil {
		http.Error(w, "Failed to create service", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(service); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
