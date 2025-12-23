package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/dataaccess"
	"github.com/rahula1008/Web_Forum/models"
)

const (
	getPostsFailedMessage    = "Failed to get all posts"
	InvalidPostIDMessage     = "Failed to read ID"
	getPostByIDFailedMessage = "Failed to get this post ID"
)

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	posts, err := dataaccess.GetAllPosts()

	if err != nil {
		sendInternalStatusServerError(
			c, getPostsFailedMessage, err,
		)
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    posts,
	})
}

func GetPostByID(c *gin.Context) {
	idStr := c.Param("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		sendBadRequestResponse(c, InvalidPostIDMessage, err)
		return
	}

	post, err := dataaccess.GetPostByID(postID)
	if err != nil {
		sendInternalStatusServerError(c, getPostByIDFailedMessage, err)
		return
	}

	c.JSON(http.StatusAccepted, Response{
		Success: true,
		Data:    post,
	})
}
