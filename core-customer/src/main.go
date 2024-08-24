package main

import (
	"core-customer/api/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	contextPath := router.Group("/core-customer")
	{
		routers.Init(contextPath)
	}

	router.Run(":9092")
}
