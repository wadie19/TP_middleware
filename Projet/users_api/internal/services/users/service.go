package users

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/users"
)

func GetAllUsers() ([]models.User, error) {
	var err error
	users, err := repository.GetAllUsers()
	if err != nil {
		logrus.Errorf("error retrieving users: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return users, nil
}

func GetUserByID(id uuid.UUID) (*models.User, error) {
	user, err := repository.GetUserByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "User not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving user: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return user, nil
}

func CreateUser(user models.User) error {
	if user.Id == nil {
		uuidCreated, err := uuid.NewV4()
		if err != nil {
			return err
		}
		user.Id = &uuidCreated
	}

	err := repository.CreateUser(user)
	if err != nil {
		logrus.Errorf("error saving user: %s", err.Error())
		return &models.CustomError{
			Message: "Failed to save user",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func UpdateUser(user models.User) error {
	// Assuming repository functions for updating users are implemented
	// Modify as per the specific logic in your repository

	err := repository.UpdateUser(user)
	if err != nil {
		logrus.Errorf("error updating user: %s", err.Error())
		return &models.CustomError{
			Message: "Failed to update user",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func DeleteUser(id uuid.UUID) error {
	err := repository.DeleteUser(id)
	if err != nil {
		logrus.Errorf("error deleting user: %s", err.Error())
		return &models.CustomError{
			Message: "Failed to delete user",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}
