package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/controllers"
	"github.com/rahula1008/Web_Forum/middleware"
)

func SetupUserRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:id", controllers.GetUserByID)
	router.GET("/users/search", controllers.SearchUserByUsername)
	router.POST("/users", controllers.CreateUser)
	router.POST("/users/signup", controllers.SignUp)
	router.POST("/users/login", controllers.Login)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
	//Only for testing purposes
	//router.GET("/users/validate", middleware.RequireAuth, controllers.Validate)
	router.GET("/users/me", middleware.RequireAuth, controllers.GetMe)
}
