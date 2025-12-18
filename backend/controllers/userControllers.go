package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/dataaccess"
	"github.com/rahula1008/Web_Forum/models"
)

const (
	getUsersFailedMessage             = "Failed to get all the users"
	InvalidIDMessage                  = "Failed to read ID"
	getUserByIDFailedMessage          = "Failed to get this user ID"
	searchUserByUsernameFailedMessage = "Failed to find users matching search"
	createUserFailedMessage           = "Failed to create user"
	updateUserFailedMessage           = "Failed to update user"
)

const (
	updateUserSuccessMessage = "Successfully updated user"
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
		sendInternalStatusServerError(c, searchUserByUsernameFailedMessage, err)
		return
	}

	c.JSON(200, Response{
		Success: true,
		Data:    users,
	})

}

func CreateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		sendBadRequestResponse(c, createUserFailedMessage, err)
		return
	}

	err = validateUser(user)

	if err != nil {
		sendBadRequestResponse(c, createUserFailedMessage, err)
		return
	}

	user.CreatedAt = time.Now()
	now := time.Now()
	user.UpdatedAt = &now

	if err = dataaccess.SaveUserToDB(&user); err != nil {
		log.Printf("DB Error creating user: %v", err)

		// Return a generic 500 server error to the client
		sendInternalStatusServerError(c, createUserFailedMessage, err)
		return
	}

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Data:    user,
		Code:    http.StatusCreated,
	})
}

func UpdateUser(c *gin.Context) {
	var updatedUser models.User

	id := c.Param("id")

	userID, err := strconv.Atoi(id)

	if err != nil {
		sendBadRequestResponse(c, updateUserFailedMessage, err)
		return
	}

	err = c.ShouldBindJSON(&updatedUser)

	if err != nil {
		sendBadRequestResponse(c, updateUserFailedMessage, err)
		return
	}

	updatedUser.ID = userID

	err = validateUser(updatedUser)

	if err != nil {
		sendBadRequestResponse(c, updateUserFailedMessage, err)
		return
	}

	err = dataaccess.UpdateUser(&updatedUser)
	if err != nil {
		sendInternalStatusServerError(c, updateUserFailedMessage, err)
		return
	}
	SendStatusOKResponse(c, updateUserSuccessMessage)
}

func validateUser(user models.User) error {
	if user.Username == "" {
		return errors.New("username cannot be blank")
	}
	if len(user.Username) > 30 {
		return errors.New("username must be at most 30 characters")
	}
	if user.Email == "" {
		return errors.New("email cannot be blank")
	}
	if len(user.Email) > 150 {
		return errors.New("email cannot be more than 150 characters")
	}
	return nil
}
