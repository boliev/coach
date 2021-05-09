package user

import (
	"github.com/boliev/coach/internal/domain"
	"github.com/boliev/coach/internal/repository"
	"github.com/boliev/coach/internal/request"
	"github.com/boliev/coach/pkg/password"
)

type UserService struct {
	repository repository.UserRepository
}

func CreateUserService(repository repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (u UserService) Create(request *request.UserCreation) (*domain.User, error) {
	password, err := password.Hash(request.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:    request.Email,
		Password: password,
	}
	_, err = u.repository.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
