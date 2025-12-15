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

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Code    int         `json:"code,omitempty"`
}

const (
	getTopicsFailedMessage             = "Failed to get all the topics"
	getTopicByIDFailedMessage          = "Failed to get this topic ID"
	searchTopicByTitleFailedMessage    = "Failed to search for this title"
	saveTopicFailedMessage             = "Failed to save topic"
	failedToReadIDMessage              = "Failed to parse ID"
	incorrectBodyForUpdateTopicMessage = "Incorrect body to update topic"
)

func GetAllTopics(c *gin.Context) {
	var topics []models.Topic

	topics, err := dataaccess.GetAllTopics()

	if err != nil {
		sendInternalStatusServerError(
			c, getTopicsFailedMessage, err,
		)
	}
	c.JSON(200, Response{
		Success: true,
		Data:    topics,
	})

}

func GetTopicByID(c *gin.Context) {

	id := c.Param("id")
	topicID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	topic, err := dataaccess.GetTopicByID(topicID)

	if err != nil {
		sendInternalStatusServerError(c, getTopicByIDFailedMessage, err)
		return
	}
	c.JSON(200, Response{
		Success: true,
		Data:    topic,
	})

}

func SearchTopic(c *gin.Context) {
	searchString := c.Query("title")
	topics, err := dataaccess.SearchTopic(searchString)

	if err != nil {
		sendInternalStatusServerError(c, searchTopicByTitleFailedMessage, err)
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
		sendInternalStatusServerError(c, saveTopicFailedMessage, err)
		return
	}

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Data:    topic,
		Message: "Topic created successfully",
		Code:    http.StatusCreated,
	})

}

func UpdateTopic(c *gin.Context) {
	id := c.Param("id")
	topicID, err := strconv.Atoi(id)

	if err != nil {
		sendBadRequestResponse(c, failedToReadIDMessage, err)
		return
	}

	var updatedTopic models.Topic
	err = c.ShouldBindJSON(&updatedTopic)

	if err != nil {
		sendBadRequestResponse(c, incorrectBodyForUpdateTopicMessage, err)
		return
	}

	updatedTopic.ID = topicID

	err = dataaccess.UpdateTopic(&updatedTopic)

	if err != nil {
		sendInternalStatusServerError(c, "Failed to update topic", err)
	}

}

func DeleteTopic(c *gin.Context) {
	id := c.Param("id")

	topicID, err := strconv.Atoi(id)

	if err != nil {
		sendBadRequestResponse(c, failedToReadIDMessage, err)
		return
	}

	err = dataaccess.DeleteTopic(topicID)

	if err != nil {
		sendInternalStatusServerError(c, "Failed to delete topic", err)
		return
	}

	c.JSON(http.StatusNoContent, Response{
		Success: true,
		Message: "Topic deleted successfully",
	})
}

func sendBadRequestResponse(c *gin.Context, message string, err error) {
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Message: message,
		Error:   err.Error(),
		Code:    http.StatusBadRequest,
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

func sendInternalStatusServerError(c *gin.Context, message string, err error) {
	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Message: message,
		Error:   err.Error(),
		Code:    http.StatusInternalServerError,
	})
}
