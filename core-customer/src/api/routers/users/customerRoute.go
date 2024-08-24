package customersRoute

import (
	"core-customer/api/controllers"
	"core-customer/api/dto/in"
	repositories "core-customer/api/infra/repositories/impl"
	"core-customer/domain/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Init(c *gin.RouterGroup, db *sqlx.DB) {

	c.POST("/customers", func(c *gin.Context) { CreateCustomer(c, db) })
	c.GET("/users", GetCustomers)
}

func CreateCustomer(c *gin.Context, db *sqlx.DB) {
	var customerIn in.CustomerInDTO

	err := c.ShouldBindJSON(&customerIn)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	}

	customerController := controllers.NewCustomerController(services.NewcustomerService(repositories.NewCustomerRepository(db)))
	customerController.CreateCustomer(customerIn)
	c.JSON(201, gin.H{
		"message": "Customer created",
	})
}

func GetCustomers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from core-customer",
	})
}
