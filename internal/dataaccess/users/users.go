package users

import (
	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func List(db *database.Database) ([]models.User, error) {
	var users []models.User
	rows, err := db.Query("SELECT `id`, `name` FROM `user`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetByName(db *database.Database, name string) (*models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT `id`, `name` FROM `user` WHERE `name` = ?", name).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(db *database.Database, user *models.User) (*models.User, error) {
	result, err := db.Exec("INSERT INTO `user` (`name`) VALUES (?)", user.Name)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(id)
	return user, nil
}

func UpdateUser(db *database.Database, name string, user *models.User) error {
	_, err := db.Exec("UPDATE `user` SET `name` = ? WHERE `name` = ?", user.Name, name)
	return err
}

func DeleteUser(db *database.Database, name string) error {
	_, err := db.Exec("DELETE FROM `user` WHERE `name` = ?", name)
	return err
}
