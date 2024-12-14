package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Error definitions
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
)

// Request/Response types
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	IsBusinessOwner bool   `json:"is_business_owner"`
}

type User struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	IsBusinessOwner bool   `json:"is_business_owner"`
}

type AuthResponse struct {
	User User `json:"user"`
}


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePasswords(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
