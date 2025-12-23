package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/dataaccess"
	"github.com/rahula1008/Web_Forum/models"
)

const (
	getCommentsFailedMessage       = "Failed to get comments"
	invalidCommentIDMessage        = "Failed to read ID"
	invalidPostIDMessage           = "Failed to read ID"
	getCommentByIDFailedMessage    = "Failed to get this comment ID"
	getCommentsByPostFailedMessage = "Failed to get comments for post"
	searchCommentsFailedMessage    = "Failed to find comments matching search"
	createCommentFailedMessage     = "Failed to create comment"
	updateCommentFailedMessage     = "Failed to update comment"
	deleteCommentFailedMessage     = "Failed to delete comment"

	updateCommentSuccessMessage = "Successfully updated comment"
)

func GetAllComments(c *gin.Context) {
	comments, err := dataaccess.GetAllComments()
	if err != nil {
		sendInternalStatusServerError(c, getCommentsFailedMessage, err)
		return
	}

	SendStatusOKResponse(c, comments)
}

func GetCommentByID(c *gin.Context) {
	idStr := c.Param("id")
	commentID, err := strconv.Atoi(idStr)
	if err != nil {
		sendBadRequestResponse(c, invalidCommentIDMessage, err)
		return
	}

	comment, err := dataaccess.GetCommentByID(commentID)
	if err != nil {
		sendInternalStatusServerError(c, getCommentByIDFailedMessage, err)
		return
	}

	SendStatusOKResponse(c, comment)
}

// GET /posts/:postId/comments
func GetCommentsByPostID(c *gin.Context) {
	postIDStr := c.Param("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		sendBadRequestResponse(c, invalidPostIDMessage, err)
		return
	}

	comments, err := dataaccess.GetCommentsByPostID(postID)
	if err != nil {
		sendInternalStatusServerError(c, getCommentsByPostFailedMessage, err)
		return
	}

	SendStatusOKResponse(c, comments)
}

func CreateComment(c *gin.Context) {
	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		sendBadRequestResponse(c, createCommentFailedMessage, err)
		return
	}

	if err := models.ValidateComment(comment); err != nil {
		sendBadRequestResponse(c, createCommentFailedMessage, err)
		return
	}

	if err := dataaccess.SaveCommentToDB(&comment); err != nil {
		sendInternalStatusServerError(c, createCommentFailedMessage, err)
		return
	}

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Data:    comment,
		Code:    http.StatusCreated,
	})
}

func UpdateComment(c *gin.Context) {
	var updated models.Comment

	idStr := c.Param("id")
	commentID, err := strconv.Atoi(idStr)
	if err != nil {
		sendBadRequestResponse(c, updateCommentFailedMessage, err)
		return
	}

	if err := c.ShouldBindJSON(&updated); err != nil {
		sendBadRequestResponse(c, updateCommentFailedMessage, err)
		return
	}

	updated.ID = commentID

	if err := models.ValidateComment(updated); err != nil {
		sendBadRequestResponse(c, updateCommentFailedMessage, err)
		return
	}

	if err := dataaccess.UpdateComment(&updated); err != nil {
		sendInternalStatusServerError(c, updateCommentFailedMessage, err)
		return
	}

	SendStatusOKResponse(c, updateCommentSuccessMessage)
}

func DeleteComment(c *gin.Context) {
	idStr := c.Param("id")
	commentID, err := strconv.Atoi(idStr)
	if err != nil {
		sendBadRequestResponse(c, deleteCommentFailedMessage, err)
		return
	}

	if err := dataaccess.DeleteComment(commentID); err != nil {
		sendInternalStatusServerError(c, deleteCommentFailedMessage, err)
		return
	}

	SendStatusNoContent(c)
}
