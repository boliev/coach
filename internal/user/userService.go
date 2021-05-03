package user

import (
	"github.com/boliev/coach/internal/domain"
	"github.com/boliev/coach/internal/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func (u UserService) Create(user domain.User) {
	u.repository.Create(user)
}
