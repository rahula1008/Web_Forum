package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/controllers"
)

func SetupTopicRoutes(router *gin.Engine) {
	router.GET("/topics", controllers.GetAllTopics)
	router.GET("/topics/:id", controllers.GetTopicByID)
	router.GET("/topics/search", controllers.SearchTopic)
	router.POST("/topics", controllers.CreateTopic)
	router.PUT("/topics/:id", controllers.UpdateTopic)
	router.DELETE("/topics/:id", controllers.DeleteTopic)
}
