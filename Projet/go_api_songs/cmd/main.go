package main

import (
	songs "middleware/example/internal/controllers/songs"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/songs", func(r chi.Router) {
		r.Get("/", songs.GetSongs)
		r.Post("/", songs.AjouterSong)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(songs.Ctx)
			r.Get("/", songs.GetSong)
			r.Put("/", songs.ModifierSong)
			r.Delete("/", songs.SupprimerSong)
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS songs (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			titre VARCHAR(255) NOT NULL,
			artiste VARCHAR(255) NOT NULL,
			description VARCHAR(255) NOT NULL,
			duree DURATION NOT NULL,
			release_date DATE NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
