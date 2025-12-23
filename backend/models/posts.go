package models

import (
	"errors"
	"time"
)

type Post struct {
	ID        int        `db:"id" json:"id"`
	Title     string     `db:"title" json:"title"`
	Body      string     `db:"body" json:"body"`
	TopicID   int        `db:"topic_id" json:"topic_id"`
	CreatorID int        `db:"creator_id" json:"creator_id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

func ValidatePost(post Post) error {
	if post.Title == "" {
		return errors.New("title cannot be blank")
	}
	if len(post.Title) > 200 {
		return errors.New("title must be at most 200 characters")
	}
	if post.Body == "" {
		return errors.New("body cannot be blank")
	}
	if post.TopicID <= 0 {
		return errors.New("topic_id must be valid")
	}
	if post.CreatorID <= 0 {
		return errors.New("creator_id must be valid")
	}
	return nil
}
