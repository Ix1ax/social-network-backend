package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid or expired token")
)

type TokenManager interface {
	GenerateToken(userID uuid.UUID) (string, error)
	ParseToken(token string) (uuid.UUID, error)
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

func (m *jwtManager) ParseToken(tokenString string) (uuid.UUID, error) {

	var claims jwt.RegisteredClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		return m.secretKey, nil
	})

	if err != nil || token == nil || !token.Valid {
		return uuid.Nil, ErrInvalidToken
	}

	return uuid.Parse(claims.Subject)

}
