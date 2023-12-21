package songs

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.Id, &data.Titre, &data.Artiste, &data.Description, &data.Duree, &data.Release_date)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return songs, err
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	var song models.Song
	err = row.Scan(&song.Id, &song.Titre, &song.Artiste, &song.Description, &song.Duree, &song.Release_date)
	if err != nil {
		return nil, err
	}
	return &song, err
}

// Ajoutersong ajoute une nouvelle song à la base de données.
func AjouterSong(id uuid.UUID, nouvelleSong *models.Song) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO songs (id, titre, artiste, description, duree, release_date) VALUES (?, ?, ?, ?, ?, ?)", id, nouvelleSong.Titre, nouvelleSong.Artiste, nouvelleSong.Description, nouvelleSong.Release_date, nouvelleSong.Release_date)
	if err != nil {
		return err
	}

	return nil
}

// Modifiersong modifie une song existante dans la base de données.
func ModifierSong(id uuid.UUID, nouvelleSong *models.Song) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE songs SET titre = ?, artiste = ?, description = ?, duree = ?, release_date = ?  WHERE id = ?", nouvelleSong.Titre, nouvelleSong.Artiste, nouvelleSong.Description, nouvelleSong.Duree, nouvelleSong.Release_date, id.String())
	if err != nil {
		return err
	}

	return nil
}

// Supprimersong supprime une song existante de la base de données.
func SupprimerSong(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM songs WHERE id = ?", id.String())
	if err != nil {
		return err
	}

	return nil
}
