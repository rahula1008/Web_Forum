package dataaccess

import (
	"errors"
	"fmt"
	"log"

	"github.com/rahula1008/Web_Forum/initializers"
	"github.com/rahula1008/Web_Forum/models"
)

func SaveTopicToDB(topic *models.Topic) error {

	insertQuery := `INSERT INTO topics (title, description, creator_id) 
		VALUES (:title, :description, :creator_id)
		returning id`

	rows, err := initializers.DB.NamedQuery(insertQuery, topic)

	if err != nil {
		log.Printf("Repository Error: Failed to save topic: %v", err)
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		return errors.New("save topic: no id returned")
	}

	if err := rows.Scan(&topic.ID); err != nil {
		return err
	}

	// If successful, the topic struct now has the new ID.
	return rows.Err()
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

func GetTopicByID(id int) (*models.Topic, error) {
	var topic models.Topic

	getTopicByIDQuery := `SELECT * from topics where id = $1`
	err := initializers.DB.Get(&topic, getTopicByIDQuery, id)

	if err != nil {
		return nil, err
	}
	return &topic, nil

}

func SearchTopic(searchString string) ([]models.Topic, error) {
	var topics []models.Topic

	findTopicLikeSearchString := `
		SELECT *
		FROM topics
		WHERE title ILIKE '%' || $1 || '%'
		ORDER BY created_at DESC
	`
	err := initializers.DB.Select(&topics, findTopicLikeSearchString, searchString)

	if err != nil {
		return nil, fmt.Errorf("search topics: %w", err)
	}
	return topics, nil

}
