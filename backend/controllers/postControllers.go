package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/dataaccess"
	"github.com/rahula1008/Web_Forum/models"
)

const (
	getPostsFailedMessage = "Failed to get all posts"
)

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	posts, err := dataaccess.GetAllPosts()

	if err != nil {
		sendInternalStatusServerError(
			c, getPostsFailedMessage, err,
		)
	}
	c.JSON(200, Response{
		Success: true,
		Data:    posts,
	})

}
