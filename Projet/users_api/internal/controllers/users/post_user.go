package users

import (
	"encoding/json"
	"net/http"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/users"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Username   string `json:"username"`
		Email      string `json:"email"`
		Premium    bool   `json:"premium"`
		Birthdate  string `json:"birthdate"`
		Country    string `json:"country"`
		// Add other necessary fields
	}

	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		logrus.Errorf("error decoding request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Generate UUID for the user
	uuidCreated, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("error generating UUID: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create the user
	newUser := models.User{
		Id:        &uuidCreated,
		Username:  requestBody.Username,
		Email:     requestBody.Email,
		Premium:   requestBody.Premium,
		Birthdate: requestBody.Birthdate,
		Country:   requestBody.Country,
	}

	// Save the new user
	err = users.CreateUser(newUser)
	if err != nil {
		logrus.Errorf("error saving user: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	responseBody, _ := json.Marshal(newUser)
	_, _ = w.Write(responseBody)
}
