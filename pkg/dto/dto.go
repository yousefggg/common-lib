package dto

type ValidateTokenResponse struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Valid  bool   `json:"valid"`
}
type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Claims struct {
    UserID int64  `json:"user_id"`
    Role   string `json:"role"`
}
