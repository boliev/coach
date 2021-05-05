package response

import "github.com/boliev/coach/internal/domain"

type UsersList struct {
	Data  []*domain.User `json:"data"`
	Count int            `json:"count"`
}
