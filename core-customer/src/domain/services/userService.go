package services

import (
	"core-customer/domain/entities"
	"fmt"
)

type UserService struct {
}

func NewUserService() UserService {
	return UserService{}
}

func (c *UserService) CreateUser(name string) entities.User {
	userCreated := entities.NewUser(name)
	fmt.Println("New user created: ", userCreated)
	return userCreated
}
