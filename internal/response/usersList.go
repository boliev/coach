package response

import "github.com/boliev/coach/internal/domain"

type UsersList struct {
	Data  []*User `json:"data"`
	Count int     `json:"count"`
}

func CreateUsersListFromDomain(users []*domain.User) *UsersList {
	var responseUsers []*User
	for _, u := range users {
		responseUsers = append(responseUsers, CreateUserFromDomain(u))
	}

	return &UsersList{
		Data:  responseUsers,
		Count: len(responseUsers),
	}
}
