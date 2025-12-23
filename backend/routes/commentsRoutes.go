package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/controllers"
)

func SetupCommentRoutes(router *gin.Engine) {
	router.GET("/comments", controllers.GetAllComments)
	router.GET("/comments/:id", controllers.GetCommentByID)
	// GET /posts/:postId/comments
	router.GET("/posts/:id/comments", controllers.GetCommentsByPostID)
	router.POST("/comments", controllers.CreateComment)
	router.PUT("/comments/:id", controllers.UpdateComment)
	router.DELETE("/comments/:id", controllers.DeleteComment)
}
