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
	c.POST("/customers/register", func(c *gin.Context) { CreateCustomer(c, db) })
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
	err = customerController.CreateCustomer(customerIn)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(201, gin.H{
		"message": "Customer created",
	})
}
