package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/dataaccess"
	"github.com/rahula1008/Web_Forum/models"
)

const (
	getUsersFailedMessage = "Failed to get all the users"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User

	users, err := dataaccess.GetAllUsers()

	if err != nil {
		sendInternalStatusServerError(
			c, getUsersFailedMessage, err,
		)
	}
	c.JSON(200, Response{
		Success: true,
		Data:    users,
	})
}
