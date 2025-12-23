package models

import (
	"errors"
	"time"
)

type Comment struct {
	ID        int        `db:"id" json:"id"`
	Body      string     `db:"body" json:"body"`
	PostID    int        `db:"post_id" json:"post_id"`
	CreatorID int        `db:"creator_id" json:"creator_id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

func ValidateComment(comment Comment) error {
	if comment.Body == "" {
		return errors.New("body cannot be blank")
	}
	if comment.PostID <= 0 {
		return errors.New("post_id must be valid")
	}
	if comment.CreatorID <= 0 {
		return errors.New("creator_id must be valid")
	}
	return nil
}
