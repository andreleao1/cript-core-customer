package usersRoute

import (
	"core-customer/api/controllers"
	repositories "core-customer/api/infra/repositories/impl"
	"core-customer/domain/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Init(c *gin.RouterGroup, db *sqlx.DB) {

	c.GET("/usersa", func(c *gin.Context) { CreateUser(c, db) })
	c.GET("/users", GetUsers)
}

func CreateUser(c *gin.Context, db *sqlx.DB) {
	userController := controllers.NewUserController(services.NewUserService(repositories.NewUserRepository(db)))
	userController.CreateUser("John Doe")
	c.JSON(201, gin.H{
		"message": "User created",
	})
}

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from core-customer",
	})
}
