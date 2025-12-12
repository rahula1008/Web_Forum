package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
