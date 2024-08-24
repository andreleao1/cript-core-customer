package usersRoute

import (
	"core-customer/api/controllers"
	"core-customer/domain/services"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.RouterGroup) {

	c.GET("/usersa", CreateUser)
	c.GET("/users", GetUsers)
}

func CreateUser(c *gin.Context) {
	userController := controllers.NewUserController(services.NewUserService())
	c.JSON(201, gin.H{
		"message": userController.CreateUser("John Doe"),
	})
}

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from core-customer",
	})
}
