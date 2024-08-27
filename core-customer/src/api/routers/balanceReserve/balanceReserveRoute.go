package balancerReserveRoute

import (
	"core-customer/core-customer/src/api/controllers"
	"core-customer/core-customer/src/api/dto/in"
	repositories "core-customer/core-customer/src/api/infra/repositories/impl"
	"core-customer/core-customer/src/domain/entities"
	"core-customer/core-customer/src/domain/services"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const basePath = "/reserves"

const (
	INTERNAL_SERVER_ERROR = "Internal server error, please try again. If the problem persists, contact support."
)

func Init(c *gin.RouterGroup, db *sqlx.DB) {
	c.POST(basePath, func(c *gin.Context) { ReserveBalance(c, db) })
	c.PATCH(basePath+"/effect/:id", func(c *gin.Context) { EffectReserve(c, db) })
	c.PATCH(basePath+"/cancel/:id", func(c *gin.Context) { CancelReserve(c, db) })
}

func ReserveBalance(c *gin.Context, db *sqlx.DB) {
	var balanceReserveIn in.ReserveBalanceInDTO

	transaction, err := db.Beginx()

	if err != nil {
		slog.Error("Error starting transaction: " + err.Error())
		c.JSON(500, gin.H{
			"error": INTERNAL_SERVER_ERROR,
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
			"message": INTERNAL_SERVER_ERROR,
		})

		return
	}

	transaction.Commit()

	c.JSON(201, gin.H{
		"message": "Balance reserved successfully",
	})
}

func EffectReserve(c *gin.Context, db *sqlx.DB) {
	reserveId := c.Param("id")

	transaction, err := db.Beginx()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	balanceReserveRepository := repositories.NewBalanceReserveRepository(transaction)
	balanceReserveService := services.NewBalanceReservationService(*balanceReserveRepository)
	balanceReserveController := controllers.NewBalanceReserveController(balanceReserveService)

	err = balanceReserveController.EffectReserve(reserveId)

	if err != nil {
		slog.Error("Error effecting reserve, appling transaction rollback: " + err.Error())
		transaction.Rollback()
		c.JSON(500, gin.H{
			"message": INTERNAL_SERVER_ERROR,
		})

		return
	}

	transaction.Commit()

	c.JSON(200, gin.H{
		"message": "Reserve effected successfully",
	})
}

func CancelReserve(c *gin.Context, db *sqlx.DB) {
	reserveId := c.Param("id")

	transaction, err := db.Beginx()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	balanceReserveRepository := repositories.NewBalanceReserveRepository(transaction)
	balanceReserveService := services.NewBalanceReservationService(*balanceReserveRepository)
	balanceReserveController := controllers.NewBalanceReserveController(balanceReserveService)

	err = balanceReserveController.CancelReserve(reserveId)

	if err != nil {
		slog.Error("Error canceling reserve, appling transaction rollback: " + err.Error())
		transaction.Rollback()
		c.JSON(500, gin.H{
			"message": INTERNAL_SERVER_ERROR,
		})

		return
	}

	transaction.Commit()

	c.JSON(200, gin.H{
		"message": "Reserve canceled successfully",
	})
}
