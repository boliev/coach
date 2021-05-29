package user

import (
	"fmt"
	"time"

	"github.com/boliev/coach/internal/domain"
	"github.com/boliev/coach/pkg/jwt_client"
	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	Secret    string
	TokenDays int
	Client    jwt_client.JwtClient
}

func NewJwtService(secret string, tokenDays int, jwtClient jwt_client.JwtClient) *JwtService {
	return &JwtService{
		Secret:    secret,
		TokenDays: tokenDays,
		Client:    jwtClient,
	}
}

func (j JwtService) Create(id string) (*domain.UserAuth, error) {
	expiresAt := time.Now().UTC().AddDate(0, 0, j.TokenDays)
	claims := &jwt.MapClaims{
		"id":        id,
		"expiresAt": expiresAt.Unix(),
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

func (j JwtService) Parse(tokenString string) (string, error) {
	token, err := j.Client.Parse(tokenString)

	if err != nil {
		return "", err
	}

	tm := time.Unix(int64(token.Claims["expiresAt"].(float64)), 0)

	if tm.Before(time.Now()) {
		return "", fmt.Errorf("token expired")
	}

	if token.Valid {
		return token.Claims["id"].(string), nil
	} else {
		return "", fmt.Errorf("token is invalid")
	}
}
