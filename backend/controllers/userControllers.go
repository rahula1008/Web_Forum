package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/dataaccess"
	"github.com/rahula1008/Web_Forum/models"
)

const (
	getUsersFailedMessage      = "Failed to get all the users"
	InvalidIDMessage           = "Failed to read ID"
	getUserByIDFailedMessage   = "Failed to get this user ID"
	searchUserByUsernameFailed = "Failed to find users matching search"
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

func GetUserByID(c *gin.Context) {

	id := c.Param("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		sendBadRequestResponse(c, InvalidIDMessage, err)
		return
	}

	user, err := dataaccess.GetUserByID(userID)

	if err != nil {
		sendInternalStatusServerError(c, getUserByIDFailedMessage, err)
		return
	}

	c.JSON(http.StatusAccepted, Response{
		Success: true,
		Data:    user,
	})

}

func SearchUserByUsername(c *gin.Context) {
	searchUsername := c.Query("username")

	users, err := dataaccess.SearchUserByUsername(searchUsername)

	if err != nil {
		sendInternalStatusServerError(c, searchUserByUsernameFailed, err)
	}

	c.JSON(200, Response{
		Success: true,
		Data:    users,
	})

}
