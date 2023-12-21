package models

import (
	"github.com/gofrs/uuid"
)

type Song struct {
	Id           *uuid.UUID `json:"id"`
	Titre        string     `json:"titre"`
	Artiste      string     `json:"artiste"`
	Description  string     `json:"description"`
	Duree        string     `json:"duree"`
	Release_date string     `json:"release_date"`
}
