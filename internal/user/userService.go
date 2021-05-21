package user

import (
	"errors"
	"fmt"

	"github.com/boliev/coach/internal/domain"
	"github.com/boliev/coach/internal/repository"
	"github.com/boliev/coach/internal/request"
	"github.com/boliev/coach/pkg/password"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	repository repository.UserRepository
	jwtCreator JwtCreator
}

func CreateUserService(repository repository.UserRepository, jwtCreator JwtCreator) *UserService {
	return &UserService{
		repository: repository,
		jwtCreator: jwtCreator,
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

func (u UserService) Auth(request *request.UserAuth) (*domain.UserAuth, error) {
	user, err := u.repository.FindByEmail(request.Email)
	if err != nil {
		logrus.Info("Authorisation failed. Wrong email: " + request.Email)
		return nil, err
	}

	check := password.Check(user.Password, request.Password)
	if !check {
		logrus.Info("Authorisation failed. Wrong password for email: " + request.Email)
		return nil, errors.New("wrong password")
	}

	userAuth, err := u.jwtCreator.Create(user.Id)
	if err != nil {
		logrus.Error(fmt.Sprintf("Authorisation failed. Cant create token: %s. %s", request.Email, err.Error()))
		return nil, err
	}

	return userAuth, nil
}
