package dataaccess

import (
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
