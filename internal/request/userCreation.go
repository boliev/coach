package request

type UserCreation struct {
	email    string `json:"name" binding:"required"`
	password string `json:"password" binding:"required"`
}
