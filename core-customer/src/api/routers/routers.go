package routers

import (
	usersRoute "core-customer/api/routers/users"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Init(c *gin.RouterGroup, db *sqlx.DB) {
	usersRoute.Init(c, db)
}
