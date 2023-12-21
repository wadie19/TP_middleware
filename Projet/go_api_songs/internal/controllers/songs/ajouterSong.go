package songs

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/songs"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// AjouterSong
// @Tags         songs
// @Summary      Ajouter une nouvelle song.
// @Description  Endpoint pour ajouter une nouvelle song.
// @Accept       json
// @Produce      json
// @Param        song body models.Song true "Nouvelle song à ajouter"
// @Success      201            {string} string
// @Failure      400            "Requête invalide"
// @Failure      500            "Erreur interne du serveur"
// @Router       /songs [post]
func AjouterSong(w http.ResponseWriter, r *http.Request) {

	// Décodez le corps de la requête JSON dans une structure de données models.Song
	var nouvelleSong models.Song
	err := json.NewDecoder(r.Body).Decode(&nouvelleSong)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Requête invalide"))
		return
	}

	// Generate UUID for the user
	uuidCreated, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("error generating UUID: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Appeler le service pour ajouter la nouvelle song
	err = songs.AjouterSong(uuidCreated, &nouvelleSong)
	if err != nil {
		logrus.Errorf("Erreur lors de l'ajout de la song : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Erreur interne du serveur"))
		return
	}
	newSong := models.Song{
		Id:           &uuidCreated,
		Titre:        *&nouvelleSong.Titre,
		Artiste:      *&nouvelleSong.Artiste,
		Description:  *&nouvelleSong.Description,
		Duree:        *&nouvelleSong.Duree,
		Release_date: *&nouvelleSong.Release_date,
	}

	// Répondre avec un statut 201 Created si tout s'est bien passé
	w.WriteHeader(http.StatusCreated)
	responseBody, _ := json.Marshal(newSong)
	_, _ = w.Write(responseBody)
}
