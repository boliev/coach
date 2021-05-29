package user_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/boliev/coach/internal/user"
	"github.com/boliev/coach/pkg/jwt_client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type JwtClientMock struct {
	jwt_client.JwtClient
	mock.Mock
}

func (j JwtClientMock) Parse(tokenString string) (*jwt_client.JwtToken, error) {
	args := j.Called(tokenString)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*jwt_client.JwtToken), args.Error(1)
}

func TestParseSuccess(t *testing.T) {
	clientMock := new(JwtClientMock)
	clientMock.On("Parse", "Some string").Return(createJwtToken("userIdString", time.Now().UTC().AddDate(0, 0, 3), true), nil).Once()
	service := user.NewJwtService("", 7, clientMock)

	userId, err := service.Parse("Some string")

	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	assert.Equal(t, userId, "userIdString")
}

func TestParseError(t *testing.T) {
	clientMock := new(JwtClientMock)
	clientMock.On("Parse", "Some string").Return(
		nil,
		fmt.Errorf("Some error"),
	).Once()

	service := user.NewJwtService("", 7, clientMock)

	userId, err := service.Parse("Some string")

	if userId != "" {
		t.Errorf("UserId should be empty when error")
	}

	assert.Equal(t, err.Error(), "Some error")
}

func TestParseTokenExpired(t *testing.T) {
	clientMock := new(JwtClientMock)
	clientMock.On("Parse", "Some string").Return(createJwtToken("userIdString", time.Now().UTC().AddDate(0, 0, -1), true), nil).Once()
	service := user.NewJwtService("", 7, clientMock)

	userId, err := service.Parse("Some string")

	if userId != "" {
		t.Errorf("UserId should be empty when error")
	}

	assert.Equal(t, err.Error(), "token expired")
}

func TestParseTokenInvalid(t *testing.T) {
	clientMock := new(JwtClientMock)
	clientMock.On("Parse", "Some string").Return(createJwtToken("userIdString", time.Now().UTC().AddDate(0, 0, 3), false), nil).Once()
	service := user.NewJwtService("", 7, clientMock)

	userId, err := service.Parse("Some string")

	if userId != "" {
		t.Errorf("UserId should be empty when error")
	}

	assert.Equal(t, err.Error(), "token is invalid")
}

func createJwtToken(id string, expiresAt time.Time, valid bool) *jwt_client.JwtToken {
	return &jwt_client.JwtToken{
		Claims: map[string]interface{}{
			"id":        id,
			"expiresAt": float64(expiresAt.Unix()),
		},
		Valid: valid,
	}
}
