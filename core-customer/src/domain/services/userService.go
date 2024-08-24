package services

import (
	"core-customer/api/infra/repositories"
	"core-customer/domain/entities"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return UserService{userRepository}
}

func (c *UserService) CreateUser(name string) {
	userCreated := entities.NewUser(name)
	c.userRepository.Create(&userCreated)
}
