package jwt_client

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type JwtClientImpl struct {
	Secret string
}

func NewJwtClientImpl(secret string) JwtClient {
	return &JwtClientImpl{
		Secret: secret,
	}
}

func (j JwtClientImpl) Parse(tokenString string) (*JwtToken, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(j.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("cant get claims from token")
	}

	return &JwtToken{
		Claims: claims,
		Valid:  token.Valid,
	}, nil
}
