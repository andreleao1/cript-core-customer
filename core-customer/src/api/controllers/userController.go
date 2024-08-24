package controllers

import (
	"core-customer/domain/entities"
	"core-customer/domain/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{userService}
}

func (u UserController) CreateUser(name string) entities.User {
	newUser := u.userService.CreateUser(name)
	return newUser
}
