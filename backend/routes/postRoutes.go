package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/controllers"
)

func SetupPostRoutes(router *gin.Engine) {
	router.GET("/posts", controllers.GetAllUsers)
	router.GET("/posts/:id", controllers.GetUserByID)
	router.GET("/posts/search", controllers.SearchUserByUsername)
	router.POST("/posts", controllers.CreateUser)
	router.PUT("/posts/:id", controllers.UpdateUser)
	router.DELETE("/posts/:id", controllers.DeleteUser)
}
