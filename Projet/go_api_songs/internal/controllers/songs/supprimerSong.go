package songs

import (
	"middleware/example/internal/models"
	"middleware/example/internal/services/songs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// Supprimersong supprime une song existante.
// @Tags         songs
// @Summary      Supprimer une song.
// @Description  Endpoint pour supprimer une song existante.
// @Param        id path string true "ID de la song à supprimer"
// @Success      204            {string} string
// @Failure      400            "ID de song invalide"
// @Failure      404            "song non trouvée"
// @Failure      500            "Erreur interne du serveur"
// @Router       /songs/{id} [delete]
func SupprimerSong(w http.ResponseWriter, r *http.Request) {

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

	// Appelez le service pour supprimer la song
	err = songs.SupprimerSong(id)
	if err != nil {
		switch err.(type) {
		case *models.CustomError:
			customError := err.(*models.CustomError)
			w.WriteHeader(customError.Code)
			_, _ = w.Write([]byte(customError.Message))
		default:
			logrus.Errorf("Erreur lors de la suppression de la song : %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Erreur interne du serveur"))
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
