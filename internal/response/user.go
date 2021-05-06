package response

import "github.com/boliev/coach/internal/domain"

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func CreateUserFromDomain(user *domain.User) *User {
	return &User{
		Id:    user.Id,
		Email: user.Email,
	}
}
