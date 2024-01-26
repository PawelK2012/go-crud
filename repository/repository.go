package repository

// https://medium.easyread.co/unit-test-sql-in-golang-5af19075e68e

import (
	"context"
	"errors"

	"github.com/PawelK2012/go-crud/models"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExist     = errors.New("row does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

// Repository implements commond DB client methods
// This allow each DB SDK to be wrapped in a Repository ie. Postgress, Redis etc
type Repository interface {
	Init(ctx context.Context) error
	GetById(ctx context.Context, id int) (models.Note, error)
	Create(ctx context.Context, note models.Note) (int64, error)
	Update(ctx context.Context, id string, note models.Note) (models.Note, error)
	GetAll(ctx context.Context) ([]models.Note, error)
}
