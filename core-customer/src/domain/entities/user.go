package entities

import "github.com/google/uuid"

type User struct {
	Id   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name" binding:"required"`
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
