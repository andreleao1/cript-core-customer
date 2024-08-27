package balancerReserveRoute

import (
	"core-customer/api/controllers"
	"core-customer/api/dto/in"
	repositories "core-customer/api/infra/repositories/impl"
	"core-customer/domain/entities"
	"core-customer/domain/services"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const basePath = "/reserves"

func Init(c *gin.RouterGroup, db *sqlx.DB) {
	c.POST(basePath, func(c *gin.Context) { ReserveBalance(c, db) })
}

func ReserveBalance(c *gin.Context, db *sqlx.DB) {
	var balanceReserveIn in.ReserveBalanceInDTO

	transaction, err := db.Beginx()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&balanceReserveIn)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	}

	balanceReserveRepository := repositories.NewBalanceReserveRepository(transaction)
	balanceReserveService := services.NewBalanceReservationService(*balanceReserveRepository)
	balanceReserveController := controllers.NewBalanceReserveController(balanceReserveService)

	balanceReserve := entities.NewBalanceReserve(balanceReserveIn.WalletId, balanceReserveIn.Amount)
	err = balanceReserveController.ReserveBalance(&balanceReserve)

	if err != nil {
		slog.Error("Error reserving balance, appling transaction rollback: " + err.Error())
		transaction.Rollback()
		c.JSON(500, gin.H{
			"message": "Error reserving balance, please try again. If the problem persists, contact support.",
		})

		return
	}

	transaction.Commit()

	c.JSON(201, gin.H{
		"message": "Balance reserved successfully",
	})
}
