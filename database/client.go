package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/PawelK2012/go-crud/models"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExist     = errors.New("row does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

// ClientInterface implements commond DB client methods
// This allow each DB SDK to be wrapped in a ClientInterface ie. Postgress, Redis etc
type ClientInterface interface {
	Init(ctx context.Context) error
	GetNoteById(ctx context.Context, id int) (*models.Note, error)
	Create(ctx context.Context, note *models.Note) (int64, error)
	GetAll(ctx context.Context) (*sql.Rows, error)
}

type Document struct {
	Id         string
	Properties map[string]interface{}
}
type ClientDBResult struct {
	Documents  []Document
	Err        error
	StatusCode int
}
