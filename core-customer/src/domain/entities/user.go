package entities

import "github.com/google/uuid"

type User struct {
	Id   uuid.UUID
	Name string
}

func NewUser(name string) User {
	return User{
		Id:   uuid.New(),
		Name: name,
	}
}

func (user User) String() string {
	return user.Name
}
