package routers

import (
	balancerReserveRoute "core-customer/core-customer/src/api/routers/balanceReserve"
	customersRoute "core-customer/core-customer/src/api/routers/users"
	walletsRoute "core-customer/core-customer/src/api/routers/wallets"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Init(c *gin.RouterGroup, db *sqlx.DB) {
	customersRoute.Init(c, db)
	walletsRoute.Init(c, db)
	balancerReserveRoute.Init(c, db)
}
