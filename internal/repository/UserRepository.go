package repository

import "github.com/boliev/coach/internal/domain"

type UserRepository interface {
	Create(user *domain.User) (interface{}, error)
	FindAll() ([]*domain.User, error)
	Find(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Delete(id string)
}
