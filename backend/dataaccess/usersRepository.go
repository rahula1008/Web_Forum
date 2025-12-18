package dataaccess

import (
	"errors"
	"log"

	"github.com/rahula1008/Web_Forum/initializers"
	"github.com/rahula1008/Web_Forum/models"
)

func GetAllUsers() ([]models.User, error) {
	users := []models.User{}
	getUsersQuery := `SELECT * FROM USERS`
	err := initializers.DB.Select(&users, getUsersQuery)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id int) (*models.User, error) {
	var user models.User

	getUserByIDQuery := `SELECT * FROM users
	WHERE ID = $1`

	err := initializers.DB.Get(&user, getUserByIDQuery, id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func SearchUserByUsername(searchUsername string) ([]models.User, error) {

	var users []models.User

	searchUserQuery :=
		`SELECT * 
	FROM users 
	WHERE username ILIKE '%' || $1 || '%'`

	err := initializers.DB.Select(&users, searchUserQuery, searchUsername)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func SaveUserToDB(user *models.User) error {
	insertQuery := `INSERT INTO users (username, email, password_hash) 
		VALUES (:username, :email, :password_hash)
		returning id`

	rows, err := initializers.DB.NamedQuery(insertQuery, user)

	if err != nil {
		log.Printf("Repository Error: Failed to save user: %v", err)
		return err
	}

	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		return errors.New("save user: no id returned")
	}

	if err := rows.Scan(&user.ID); err != nil {
		return err
	}

	// If successful, the user struct now has the new ID.
	return rows.Err()

}

func UpdateUser(user *models.User) error {
	updateQuery := `
	UPDATE users
	SET username=:username, email=:email, password_hash=:password_hash
	WHERE id = :id`

	result, err := initializers.DB.NamedExec(updateQuery, user)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("update user: no rows affected")
	}
	return nil
}
