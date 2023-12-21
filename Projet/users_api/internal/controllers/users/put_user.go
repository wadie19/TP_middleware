package users

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/services/users"
)

// UpdateUser is used to update a user.
// Assumes that the User model structure is defined.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, _ := ctx.Value("userId").(uuid.UUID)

	// Get the request data for user update
	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		logrus.Errorf("Error decoding request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Fetch the user data to update
	user, err := users.GetUserByID(userID)
	if err != nil {
		if _, ok := err.(*models.NotFoundError); ok {
			logrus.WithField("userID", userID.String()).Error("User not found")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		logrus.WithError(err).Error("Error getting user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Update all user fields with new values from the request
	// This assumes that the JSON request contains all the fields of the User model
	// Update all fields even if they are empty or zero values
	*user = updatedUser

	// Save the updated user data in the database
	err = users.UpdateUser(*user)
	if err != nil {
		logrus.WithError(err).Error("Error updating user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the updated user in response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
