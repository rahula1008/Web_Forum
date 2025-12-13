package controllers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/dataaccess"
	"github.com/rahula1008/Web_Forum/models"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Code    int         `json:"code,omitempty"`
}

func GetAllTopics(c *gin.Context) {
	var topics []models.Topic

	topics, err := dataaccess.GetAllTopics()

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to get topics",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}
	c.JSON(200, Response{
		Success: true,
		Data:    topics,
	})

}

func CreateTopic(c *gin.Context) {
	var topic models.Topic

	err := c.ShouldBindJSON(&topic)

	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	err = validateTopic(topic)
	if err != nil {
		c.JSON(400, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	topic.CreatedAt = time.Now()
	now := time.Now()
	topic.UpdatedAt = &now
	topic.CreatorID = 1 // Placeholder

	if err = dataaccess.SaveTopicToDB(&topic); err != nil {
		log.Printf("DB Error creating topic: %v", err)

		// Return a generic 500 server error to the client
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Error:   "Could not create topic due to a server error.",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Data:    topic,
		Message: "Topic created successfully",
		Code:    http.StatusCreated,
	})

}

func validateTopic(topic models.Topic) error {
	if topic.Description == "" {
		return errors.New("description cannot be blank")
	}
	if topic.Title == "" {
		return errors.New("title cannot be blank")
	}
	if len(topic.Title) > 100 {
		return errors.New("length of title must be at most 100")
	}
	return nil
}
