package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/initializers"
	"github.com/rahula1008/Web_Forum/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routes.SetupTopicRoutes(router)
	routes.SetupUserRoutes(router)
	routes.SetupPostRoutes(router)
	routes.SetupCommentRoutes(router)
	router.Run()
}
