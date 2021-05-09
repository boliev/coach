package response

import "github.com/boliev/coach/internal/domain"

type UserAuth struct {
	Token string `json:"token"`
}

func CreateUserAuthFromDomain(auth *domain.UserAuth) *UserAuth {
	return &UserAuth{
		Token: auth.Token,
	}
}
