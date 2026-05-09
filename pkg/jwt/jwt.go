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

func (m *TokenManager) ValidateToken(tokenString string) (uuid.UUID, string, error) {

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(m.secretKey), nil
    })

    if err != nil {
        return uuid.Nil, "", err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    
        sub, ok := claims["sub"].(string)
        if !ok {
            return uuid.Nil, "", errors.New("invalid subject in token")
        }

        userID, err := uuid.Parse(sub)
        if err != nil {
            return uuid.Nil, "", errors.New("invalid user id format")
        }

        role, ok := claims["role"].(string)
        if !ok {
            return uuid.Nil, "", errors.New("role not found in token")
        }

        return userID, role, nil
    }

    return uuid.Nil, "", errors.New("invalid token claims")
}