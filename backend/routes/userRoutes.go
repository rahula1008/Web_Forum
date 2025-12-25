package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/controllers"
)

func SetupUserRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:id", controllers.GetUserByID)
	router.GET("/users/search", controllers.SearchUserByUsername)
	router.POST("/users", controllers.CreateUser)
	router.POST("/users/signup", controllers.SignUp)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
}
