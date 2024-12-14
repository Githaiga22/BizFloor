package handlers

import (
	"encoding/json"
	"net/http"

	"group7/auth"
	"group7/models"

	"gorm.io/gorm"
)

type ProfileHandler struct {
	DB *gorm.DB
}

func (h *ProfileHandler) profileHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	session := r.Context().Value("session").(*auth.Session)

	if session == nil {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	// Fetch user from database to get is_business_owner status
	var user models.User
	if err := h.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := auth.User{
		ID:              userID,
		Name:            session.UserName,
		Email:           session.UserEmail,
		IsBusinessOwner: user.IsBusinessOwner,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
