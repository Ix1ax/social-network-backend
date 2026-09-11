package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenManager interface {
	GenerateToken(userID uuid.UUID) (string, error)
}

type jwtManager struct {
	secretKey []byte
	ttl       time.Duration
}

func NewJwtManager(secret string, ttl time.Duration) TokenManager {
	return &jwtManager{
		secretKey: []byte(secret),
		ttl:       ttl,
	}
}

func (m *jwtManager) GenerateToken(userID uuid.UUID) (string, error) {

	claims := jwt.RegisteredClaims{
		Subject:   userID.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.ttl)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.secretKey)

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
