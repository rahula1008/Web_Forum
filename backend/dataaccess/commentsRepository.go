package dataaccess

import (
	"errors"
	"fmt"

	"github.com/rahula1008/Web_Forum/initializers"
	"github.com/rahula1008/Web_Forum/models"
)

func GetAllComments() ([]models.Comment, error) {
	comments := []models.Comment{}
	q := `SELECT * FROM comments ORDER BY created_at DESC`
	if err := initializers.DB.Select(&comments, q); err != nil {
		return nil, err
	}
	return comments, nil
}

func GetCommentByID(id int) (*models.Comment, error) {
	var comment models.Comment
	q := `SELECT * FROM comments WHERE id = $1`
	if err := initializers.DB.Get(&comment, q, id); err != nil {
		return nil, err
	}
	return &comment, nil
}

func GetCommentsByPostID(postId int) ([]models.Comment, error) {
	var comments []models.Comment
	searchQuery := `
		SELECT * FROM comments
		WHERE post_id = $1
	`
	err := initializers.DB.Select(&comments, searchQuery, postId)

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func SaveCommentToDB(comment *models.Comment) error {
	insertQuery := `
		INSERT INTO comments (body, post_id, creator_id)
		VALUES (:body, :post_id, :creator_id)
		RETURNING id
	`

	rows, err := initializers.DB.NamedQuery(insertQuery, comment)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		return errors.New("save comment: no id returned")
	}

	if err := rows.Scan(&comment.ID); err != nil {
		return err
	}

	return rows.Err()
}

func UpdateComment(comment *models.Comment) error {
	updateQuery := `
		UPDATE comments
		SET body = :body
		WHERE id = :id
	`

	result, err := initializers.DB.NamedExec(updateQuery, comment)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("update comment: no rows affected")
	}

	return nil
}

func DeleteComment(id int) error {
	deleteQuery := `DELETE FROM comments WHERE id = $1`

	result, err := initializers.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("delete comment: no rows affected")
	}

	return nil
}
