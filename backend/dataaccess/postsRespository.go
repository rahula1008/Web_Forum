package dataaccess

import (
	"fmt"

	"github.com/rahula1008/Web_Forum/initializers"
	"github.com/rahula1008/Web_Forum/models"
)

func GetAllPosts() ([]models.Post, error) {
	posts := []models.Post{}
	getPostsQuery := `SELECT * FROM posts ORDER BY created_at DESC`
	if err := initializers.DB.Select(&posts, getPostsQuery); err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPostByID(id int) (*models.Post, error) {
	var post models.Post
	getPostByIDQuery := `SELECT * FROM posts WHERE id = $1`
	if err := initializers.DB.Get(&post, getPostByIDQuery, id); err != nil {
		return nil, err
	}
	return &post, nil
}

func SearchPost(searchString string) ([]models.Post, error) {
	var posts []models.Post
	q := `
		SELECT *
		FROM posts
		WHERE title ILIKE '%' || $1 || '%'
		ORDER BY created_at DESC
	`
	if err := initializers.DB.Select(&posts, q, searchString); err != nil {
		return nil, fmt.Errorf("search posts: %w", err)
	}
	return posts, nil
}
