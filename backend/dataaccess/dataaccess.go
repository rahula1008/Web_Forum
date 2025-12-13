package dataaccess

import (
	"log"

	"github.com/rahula1008/Web_Forum/initializers"
	"github.com/rahula1008/Web_Forum/models"
)

func SaveTopicToDB(topic *models.Topic) error {

	insertQuery := `INSERT INTO topics (title, description, creator_id, created_at, updated_at) 
		VALUES (:title, :description, :creator_id, :created_at, :updated_at)
		returning id`

	rows, err := initializers.DB.NamedQuery(insertQuery, topic)

	if err != nil {
		log.Printf("Repository Error: Failed to save topic: %v", err)
		return err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&topic.ID); err != nil {
			return err
		}
	}

	// If successful, the topic struct now has the new ID.
	return nil
}

func GetAllTopics() ([]models.Topic, error) {
	topics := []models.Topic{}
	getTopicsQuery := `SELECT * FROM TOPICS`
	err := initializers.DB.Select(&topics, getTopicsQuery)
	if err != nil {
		return nil, err
	}

	return topics, nil
}
