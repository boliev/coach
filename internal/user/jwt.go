package user

import (
	"time"

	"github.com/boliev/coach/internal/domain"
	"github.com/dgrijalva/jwt-go"
)

type JwtCreator struct {
	Secret    string
	TokenDays int
}

func NewJwtCreator(secret string, tokenDays int) *JwtCreator {
	return &JwtCreator{
		Secret:    secret,
		TokenDays: tokenDays,
	}
}

func (j JwtCreator) Create(id string) (*domain.UserAuth, error) {
	expiresAt := time.Now().UTC().AddDate(0, 0, j.TokenDays)
	claims := &jwt.MapClaims{
		"Id":        id,
		"expiresAt": expiresAt,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.Secret))
	if err != nil {
		return nil, err
	}

	return &domain.UserAuth{
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}
