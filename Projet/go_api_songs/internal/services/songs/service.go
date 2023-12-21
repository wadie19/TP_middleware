package songs

import (
	"database/sql"
	"errors"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/songs"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllSongs() ([]models.Song, error) {
	var err error
	// calling repository
	songs, err := repository.GetAllSongs()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return songs, nil
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	song, err := repository.GetSongById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "song not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}

// Ajoutersong ajoute une nouvelle song à la base de données.
func AjouterSong(id uuid.UUID, nouvelleSong *models.Song) error {
	err := repository.AjouterSong(id, nouvelleSong)
	if err != nil {
		logrus.Errorf("Erreur lors de l'ajout de la song : %s", err.Error())

		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

// Modifiersong modifie une song existante.
func ModifierSong(id uuid.UUID, nouvelleSong *models.Song) error {
	// Ajoutez ici toute logique métier nécessaire avant d'appeler le repository.

	// Appelez la méthode correspondante dans le repository pour effectuer la modification.
	err := repository.ModifierSong(id, nouvelleSong)
	if err != nil {
		// Gérez les erreurs selon vos besoins.
		return err
	}

	// Ajoutez ici toute logique métier nécessaire après la modification.

	return nil
}

func SupprimerSong(id uuid.UUID) error {

	// Appelez la méthode correspondante dans le repository pour effectuer la suppression.
	err := repository.SupprimerSong(id)
	if err != nil {
		logrus.Errorf("error deleting song: %s", err.Error())
		return &models.CustomError{
			Message: "Failed to delete song",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}
