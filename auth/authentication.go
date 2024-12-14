package auth

import (
	"encoding/json"
	"net/http"

	"group7/models"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db           *gorm.DB
	sessionStore *SessionStore
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		db:           db,
		sessionStore: NewSessionStore(),
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := h.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	session, err := h.sessionStore.CreateSession(
		user.ID,
		user.Name,
		user.Email,
		user.IsBusinessOwner,
		r.RemoteAddr,
	)
	if err != nil {
		http.Error(w, "Error creating session", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    session.ID.String(),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400, 
	})

	response := AuthResponse{
		User: User{
			ID:              user.ID,
			Name:            user.Name,
			Email:           user.Email,
			Phone:           user.Phone,
			IsBusinessOwner: user.IsBusinessOwner,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		sessionID, err := uuid.FromString(cookie.Value)
		if err == nil {
			h.sessionStore.DeleteSession(sessionID)
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Email == "" || req.Password == "" || req.Name == "" || req.Phone == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	// Check if user already exists
	var existingUser models.User
	if err := h.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		http.Error(w, ErrUserExists.Error(), http.StatusConflict)
		return
	}

	// Hash password
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Error processing password", http.StatusInternalServerError)
		return
	}

	// Create new user with IsBusinessOwner field
	newUser := models.User{
		Name:            req.Name,
		Email:           req.Email,
		Phone:           req.Phone,
		PasswordHash:    hashedPassword,
		IsBusinessOwner: req.IsBusinessOwner, // Set the business owner status
	}

	if err := h.db.Create(&newUser).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Create session
	session, err := h.sessionStore.CreateSession(
		newUser.ID,
		newUser.Name,
		newUser.Email,
		newUser.IsBusinessOwner,
		r.RemoteAddr,
	)
	if err != nil {
		http.Error(w, "Error creating session", http.StatusInternalServerError)
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    session.ID.String(),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400, 
	})

	// Create response
	response := AuthResponse{
		User: User{
			ID:              newUser.ID,
			Name:            newUser.Name,
			Email:           newUser.Email,
			Phone:           newUser.Phone,
			IsBusinessOwner: newUser.IsBusinessOwner,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
