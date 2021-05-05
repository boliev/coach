package request

import "github.com/boliev/coach/internal/domain"

type UserCreation struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (r UserCreation) ToDomain() *domain.User {
	return &domain.User{
		Email:    r.Email,
		Password: r.Password,
	}
}
