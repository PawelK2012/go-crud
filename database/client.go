package database

import (
	"context"

	"github.com/PawelK2012/go-crud/models"
)

// ClientInterface implements commond DB client methods
// This allow each DB SDK to be wrapped in a ClientInterface ie. Postgress, Redis etc
type ClientInterface interface {
	Init(ctx context.Context) error
	GetNoteById(ctx context.Context, id int) (*models.Note, error)
	CreateNote(ctx context.Context, note *models.Note) error
}
