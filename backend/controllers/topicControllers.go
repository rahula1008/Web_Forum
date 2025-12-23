package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahula1008/Web_Forum/dataaccess"
	"github.com/rahula1008/Web_Forum/models"
)

const (
	getTopicsFailedMessage             = "Failed to get all the topics"
	getTopicByIDFailedMessage          = "Failed to get this topic ID"
	searchTopicByTitleFailedMessage    = "Failed to search for this title"
	saveTopicFailedMessage             = "Failed to save topic"
	failedToReadIDMessage              = "Failed to parse ID"
	incorrectBodyForUpdateTopicMessage = "Incorrect body to update topic"
	updateTopicFailedMessage           = "Failed to update topic"
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

	err = models.ValidateTopic(topic)
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

	err = models.ValidateTopic(updatedTopic)

	if err != nil {
		sendBadRequestResponse(c, updateTopicFailedMessage, err)
		return
	}

	err = dataaccess.UpdateTopic(&updatedTopic)

	if err != nil {
		sendInternalStatusServerError(c, "Failed to update topic", err)
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    updatedTopic,
	})

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
