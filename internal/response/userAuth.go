package response

import (
	"time"

	"github.com/boliev/coach/internal/domain"
)

type UserAuth struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

func CreateUserAuthFromDomain(auth *domain.UserAuth) *UserAuth {
	return &UserAuth{
		Token:     auth.Token,
		ExpiresAt: auth.ExpiresAt,
	}
}
