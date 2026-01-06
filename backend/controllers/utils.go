package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Code    int         `json:"code,omitempty"`
}

type UserResponse struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func SendStatusOKResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Code:    http.StatusOK,
		Data:    data,
	})
}

func sendStatusCreatedResponseUser(c *gin.Context) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Code:    http.StatusCreated,
	})
}

func sendInternalStatusServerError(c *gin.Context, message string, err error) {
	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Message: message,
		Error:   err.Error(),
		Code:    http.StatusInternalServerError,
	})
}

func sendBadRequestResponse(c *gin.Context, message string, err error) {

	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: message,
			Error:   err.Error(),
			Code:    http.StatusBadRequest,
		})
	} else {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: message,
			Code:    http.StatusBadRequest,
		})
	}

}

func SendStatusNoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, Response{
		Success: true,
		Code:    http.StatusNoContent,
	})
}

func sendStatusUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Success: false,
		Code:    http.StatusUnauthorized,
	})
}
