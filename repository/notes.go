package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/PawelK2012/go-crud/models"
)

const (
	NOTE_TBL = "notes"
)

func (repo *Repository) CreateNote(ctx context.Context, n *models.Note) (*models.Note, error) {
	log.Println("---> entered repository:")
	id, err := repo.Postgress.Create(ctx, NOTE_TBL, n)
	if err != nil {
		return nil, err
	}
	n.Id = id
	return n, nil
}

func (repo *Repository) GetAllNotes(ctx context.Context) (*models.Note, error) {
	rows, err := repo.Postgress.GetAll(ctx, NOTE_TBL)
	if err != nil {
		return nil, err
	}
	fmt.Printf("---> retrived rows 222 %+v\n", rows)
	return nil, nil
}
