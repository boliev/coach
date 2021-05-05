package domain

type User struct {
	Id       int `json:"_id"`
	Email    string
	Password string
}
