package dto

import "github.com/google/uuid"

type ValidateTokenResponse struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	Valid  bool      `json:"valid"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
}