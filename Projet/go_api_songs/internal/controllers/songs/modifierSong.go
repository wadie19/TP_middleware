package songs

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/songs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// Modifiersong modifie une song existante.
// @Tags         songs
// @Summary      Modifier une song.
// @Description  Endpoint pour modifier une song existante.
// @Accept       json
// @Produce      json
// @Param        id path string true "ID de la song à modifier"
// @Param        song body models.Song true "Nouvelle valeur de la song"
// @Success      200            {string} string
// @Failure      400            "Requête invalide"
// @Failure      404            "song non trouvée"
// @Failure      500            "Erreur interne du serveur"
// @Router       /songs/{id} [put]
func ModifierSong(w http.ResponseWriter, r *http.Request) {
	// Récupérez l'ID de la song depuis les paramètres de l'URL
	songID := chi.URLParam(r, "id")

	// Parsez l'ID en UUID
	id, err := uuid.FromString(songID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("ID de song invalide"))
		return
	}

	// Vérifiez si la song existe
	existingSong, err := songs.GetSongById(id)
	if err != nil {
		switch err.(type) {
		case *models.CustomError:
			customError := err.(*models.CustomError)
			w.WriteHeader(customError.Code)
			_, _ = w.Write([]byte(customError.Message))
		default:
			logrus.Errorf("Erreur lors de la récupération de la song : %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Erreur interne du serveur"))
		}
		return
	}

	// Si la song n'existe pas, renvoyez une erreur 404
	if existingSong == nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Song non trouvée"))
		return
	}

	// Decode le corps de la requête JSON dans une nouvelle valeur de song
	var nouvelleSong models.Song
	err = json.NewDecoder(r.Body).Decode(&nouvelleSong)
	if err != nil {
		logrus.Errorf("Error decoding request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Requête invalide"))
		return
	}

	// Appelez le service pour modifier la song
	err = songs.ModifierSong(id, &nouvelleSong)
	if err != nil {
		switch err.(type) {
		case *models.CustomError:
			customError := err.(*models.CustomError)
			w.WriteHeader(customError.Code)
			_, _ = w.Write([]byte(customError.Message))
		default:
			logrus.Errorf("Erreur lors de la modification de la song : %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Erreur interne du serveur"))
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	responseBody, _ := json.Marshal(nouvelleSong)
	_, _ = w.Write(responseBody)

}
