package user

import "github.com/boliev/coach/internal/domain"

type JwtCreator interface {
	Create(id string) (*domain.UserAuth, error)
}
