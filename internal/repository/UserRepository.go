package repository

import "github.com/boliev/coach/internal/domain"

type UserRepository interface {
	Create(user domain.User)
}
