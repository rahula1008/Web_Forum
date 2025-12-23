package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/dataaccess"
	"github.com/rahula1008/Web_Forum/models"
)

const (
	getPostsFailedMessage      = "Failed to get all posts"
	InvalidPostIDMessage       = "Failed to read ID"
	getPostByIDFailedMessage   = "Failed to get this post ID"
	searchPostByTitleFailedMsg = "Failed to find posts matching the search"
	createPostFailedMessage    = "Failed to create post"
	updatePostFailedMessage    = "Failed to update post"
	deletePostFailedMessage    = "Failed to delete post"
)

const (
	updatePostSuccessMessage = "Successfully updated post"
)

func GetAllPosts(c *gin.Context) {
	var posts []models.Post

	posts, err := dataaccess.GetAllPosts()

	if err != nil {
		sendInternalStatusServerError(c, getPostsFailedMessage, err)
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

func SearchPostByTitle(c *gin.Context) {
	searchTitle := c.Query("title")

	posts, err := dataaccess.SearchPost(searchTitle)
	if err != nil {
		sendInternalStatusServerError(c, searchPostByTitleFailedMsg, err)
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    posts,
	})
}

func CreatePost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		sendBadRequestResponse(c, createPostFailedMessage, err)
		return
	}

	if err := validatePost(post); err != nil {
		sendBadRequestResponse(c, createPostFailedMessage, err)
		return
	}

	if err := dataaccess.SavePostToDB(&post); err != nil {
		log.Printf("DB Error creating post: %v", err)
		sendInternalStatusServerError(c, createPostFailedMessage, err)
		return
	}

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Data:    post,
		Code:    http.StatusCreated,
	})
}

func UpdatePost(c *gin.Context) {
	var updatedPost models.Post

	idStr := c.Param("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		sendBadRequestResponse(c, updatePostFailedMessage, err)
		return
	}

	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		sendBadRequestResponse(c, updatePostFailedMessage, err)
		return
	}

	updatedPost.ID = postID

	if err := validatePost(updatedPost); err != nil {
		sendBadRequestResponse(c, updatePostFailedMessage, err)
		return
	}

	if err := dataaccess.UpdatePost(&updatedPost); err != nil {
		sendInternalStatusServerError(c, updatePostFailedMessage, err)
		return
	}

	SendStatusOKResponse(c, updatePostSuccessMessage)
}

func DeletePost(c *gin.Context) {
	idStr := c.Param("id")
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		sendBadRequestResponse(c, deletePostFailedMessage, err)
		return
	}

	if err := dataaccess.DeletePost(postID); err != nil {
		sendInternalStatusServerError(c, deletePostFailedMessage, err)
		return
	}

	SendStatusNoContent(c)
}

func validatePost(post models.Post) error {
	if post.Title == "" {
		return errors.New("title cannot be blank")
	}
	if len(post.Title) > 200 {
		return errors.New("title must be at most 200 characters")
	}
	if post.Body == "" {
		return errors.New("body cannot be blank")
	}
	if post.TopicID <= 0 {
		return errors.New("topic_id must be valid")
	}
	if post.CreatorID <= 0 {
		return errors.New("creator_id must be valid")
	}
	return nil
}
