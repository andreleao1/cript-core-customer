package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()
	contextPath := router.Group("/core-customer")
	{
		contextPath.GET("/users", GetUsers)
	}

	router.Run()
}

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from core-customer",
	})
}
