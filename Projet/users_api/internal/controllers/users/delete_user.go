package users

import (
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/users"
	"net/http"
)

// DeleteUser deletes a user by ID.
// Assumes that the User model structure is defined.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userId").(uuid.UUID)
	if !ok {
		logrus.Error("Invalid UUID format")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := users.DeleteUser(userID)
	if err != nil {
		switch err.(type) {
		case *models.NotFoundError:
			w.WriteHeader(http.StatusNotFound)
		default:
			logrus.Errorf("Error deleting user: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
