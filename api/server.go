package main

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	port := "8080"
	router.Run(":" + port)
}
