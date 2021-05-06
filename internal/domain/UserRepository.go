package domain

type UserRepository interface {
	Create(user *User) (interface{}, error)
	FindAll() ([]*User, error)
	Find(id string) (*User, error)
	Delete(id string)
}
