package dataaccess

import (
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
