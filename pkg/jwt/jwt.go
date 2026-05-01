package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenManager struct {
	secretKey     string
	tokenTTL      time.Duration
}

func NewTokenManager(secretKey string, tokenTTL time.Duration) (*TokenManager, error) {
	if secretKey == "" {
		return nil, errors.New("empty secret key")
	}
	return &TokenManager{
		secretKey: secretKey,
		tokenTTL:  tokenTTL,
	}, nil
}

func (m *TokenManager) GenerateToken(userID uuid.UUID, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID.String(),
		"role": role,
		"exp":  time.Now().Add(m.tokenTTL).Unix(),
		"iat":  time.Now().Unix(),
	})

	return token.SignedString([]byte(m.secretKey))
}