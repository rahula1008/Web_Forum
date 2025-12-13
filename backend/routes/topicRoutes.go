package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/controllers"
)

func SetupTopicRoutes(router *gin.Engine) {
	router.GET("/topics", controllers.GetAllTopics)

	router.POST("/topics", controllers.CreateTopic)
}
