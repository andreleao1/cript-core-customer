package walletsRoute

import (
	"core-customer/core-customer/src/api/controllers"
	repositories "core-customer/core-customer/src/api/infra/repositories/impl"
	"core-customer/core-customer/src/domain/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Init(c *gin.RouterGroup, db *sqlx.DB) {
	c.GET("/wallets/customer/:customerId", func(c *gin.Context) { GetWalletByCustomerId(c, db) })
}

func GetWalletByCustomerId(c *gin.Context, db *sqlx.DB) {
	customerId := c.Param("customerId")

	dbTransaction, err := db.Beginx()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	walletController := controllers.NewWalletController(services.NewWalletService(repositories.NewWalletRepository(dbTransaction)))
	wallet, err := walletController.GetWalletByCustomerId(customerId)

	if err != nil {
		c.JSON(404, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, wallet)
}
