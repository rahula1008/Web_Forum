package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Code    int         `json:"code,omitempty"`
}

func SendStatusOKResponse(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Code:    http.StatusOK,
	})
}

func SendStatusNoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, Response{
		Success: true,
		Code:    http.StatusNoContent,
	})
}
