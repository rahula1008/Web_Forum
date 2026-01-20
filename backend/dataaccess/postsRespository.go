package dataaccess

import (
	"errors"
	"fmt"
	"log"

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

func SavePostToDB(post *models.Post) error {
	insertQuery := `
		INSERT INTO posts (title, body, topic_id, creator_id)
		VALUES (:title, :body, :topic_id, :creator_id)
		RETURNING id
	`

	rows, err := initializers.DB.NamedQuery(insertQuery, post)
	if err != nil {
		log.Printf("Repository Error: Failed to save post: %v", err)
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		return errors.New("save post: no id returned")
	}

	if err := rows.Scan(&post.ID); err != nil {
		return err
	}

	return rows.Err()
}

func UpdatePost(post *models.Post) error {
	updateQuery := `
		UPDATE posts
		SET title = :title,
		    body  = :body
		WHERE id = :id
	`

	result, err := initializers.DB.NamedExec(updateQuery, post)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("update post: no rows affected")
	}

	return nil
}

func DeletePost(id int) error {
	deleteQuery := `DELETE FROM posts WHERE id = $1`

	result, err := initializers.DB.Exec(deleteQuery, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if rowsAffected == 0 {
		fmt.Println(err)
		return fmt.Errorf("delete post: no rows affected")
	}

	return nil
}

func GetPostsByTopicID(topicID int) ([]models.Post, error) {
	posts := []models.Post{}
	q := `
		SELECT *
		FROM posts
		WHERE topic_id = $1
		ORDER BY created_at DESC
	`
	if err := initializers.DB.Select(&posts, q, topicID); err != nil {
		return nil, err
	}
	return posts, nil
}
