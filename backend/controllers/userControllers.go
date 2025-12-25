package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/dataaccess"
	"github.com/rahula1008/Web_Forum/models"
	"golang.org/x/crypto/bcrypt"
)

const (
	getUsersFailedMessage             = "Failed to get all the users"
	InvalidIDMessage                  = "Failed to read ID"
	getUserByIDFailedMessage          = "Failed to get this user ID"
	searchUserByUsernameFailedMessage = "Failed to find users matching search"
	createUserFailedMessage           = "Failed to create user"
	updateUserFailedMessage           = "Failed to update user"
	deleteUserFailedMessage           = "Failed to delete user"
	failedToReadBodyMessage           = "Failed to read body of user"
	failedToHashPasswordMessage       = "Failed to hash password"
	invalidUserMessage                = "Invalid fields for a user"
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

	err = models.ValidateUser(user)

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

func SignUp(c *gin.Context) {

	// Use this since the password that comes in is not hashed so it won't bind
	// to the user model (which expects password_hash)
	var body struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	//Get the username, email, and password off req body

	if err := c.ShouldBindJSON(&body); err != nil {
		sendBadRequestResponse(c, failedToReadBodyMessage, err)
		return
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		sendBadRequestResponse(c, failedToHashPasswordMessage, err)
		return
	}

	//Create the user
	user := models.User{
		Username:     body.Username,
		Email:        body.Email,
		PasswordHash: string(hash),
	}

	//Validate the user
	if err = models.ValidateUser(user); err != nil {
		sendBadRequestResponse(c, invalidUserMessage, err)
		return
	}

	//Save the user to the database
	if err = dataaccess.SaveUserToDB(&user); err != nil {
		log.Printf("DB Error creating user: %v", err)

		// Return a generic 500 server error to the client
		sendInternalStatusServerError(c, createUserFailedMessage, err)
		return
	}
	// Respond
	sendStatusCreatedResponseUser(c)
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

	err = models.ValidateUser(updatedUser)

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

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	userID, err := strconv.Atoi(id)

	if err != nil {
		sendBadRequestResponse(c, deleteUserFailedMessage, err)
		return
	}

	err = dataaccess.DeleteUser(userID)

	if err != nil {
		sendInternalStatusServerError(c, deleteUserFailedMessage, err)
		return
	}

	SendStatusNoContent(c)
}
