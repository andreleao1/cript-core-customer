package routers

import (
	customersRoute "core-customer/api/routers/users"
	walletsRoute "core-customer/api/routers/wallets"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Init(c *gin.RouterGroup, db *sqlx.DB) {
	customersRoute.Init(c, db)
	walletsRoute.Init(c, db)
}
