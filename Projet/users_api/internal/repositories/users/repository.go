package users

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	users := []models.User{}
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Premium, &user.Birthdate, &user.Country)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	_ = rows.Close()

	return users, err
}

func GetUserByID(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE id=?", id.String())
	helpers.CloseDB(db)

	var user models.User
	err = row.Scan(&user.Id, &user.Username, &user.Email, &user.Premium, &user.Birthdate, &user.Country)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func CreateUser(user models.User) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO users (id, username, email, premium, birthdate, country) VALUES (?, ?, ?, ?, ?, ?)",
		user.Id.String(), user.Username, user.Email, user.Premium, user.Birthdate, user.Country)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user models.User) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE users SET username=?, email=?, premium=?, birthdate=?, country=? WHERE id=?",
		user.Username, user.Email, user.Premium, user.Birthdate, user.Country, user.Id.String())
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM users WHERE id=?", id.String())
	if err != nil {
		return err
	}
	return nil
}
