package models

import "time"

type Post struct {
	ID        int        `db:"id" json:"id"`
	Title     string     `db:"title" json:"title"`
	Body      string     `db:"body" json:"body"`
	TopicID   int        `db:"topic_id" json:"topic_id"`
	CreatorID int        `db:"creator_id" json:"creator_id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
