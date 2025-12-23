package models

import (
	"errors"
	"time"
)

type Topic struct {
	ID          int        `db:"id" json:"id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	CreatorID   int        `db:"creator_id" json:"creator_id"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}

func ValidateTopic(topic Topic) error {
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
