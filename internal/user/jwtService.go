package user

import (
	"fmt"
	"time"

	"github.com/boliev/coach/internal/domain"
	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	Secret    string
	TokenDays int
}

func NewJwtCreator(secret string, tokenDays int) *JwtService {
	return &JwtService{
		Secret:    secret,
		TokenDays: tokenDays,
	}
}

func (j JwtService) Create(id string) (*domain.UserAuth, error) {
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

func (j JwtService) Parse(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(j.Secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["Id"].(string), nil
	} else {
		return "", err
	}
}
