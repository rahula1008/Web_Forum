package routes

import "github.com/gin-gonic/gin"

func SetupTopicRoutes(router *gin.Engine) {
	router.GET("/topics", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
