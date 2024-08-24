package routers

import (
	usersRoute "core-customer/api/routers/users"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.RouterGroup) {
	usersRoute.Init(c)
}
