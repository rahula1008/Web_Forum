package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/controllers"
)

func SetupPostRoutes(router *gin.Engine) {
	router.GET("/posts", controllers.GetAllPosts)
	router.GET("/posts/:id", controllers.GetPostByID)
	router.GET("/posts/search", controllers.SearchPostByTitle)
	// router.POST("/posts", controllers.CreatePost)
	// router.PUT("/posts/:id", controllers.UpdatePost)
	// router.DELETE("/posts/:id", controllers.DeletePost)
}
