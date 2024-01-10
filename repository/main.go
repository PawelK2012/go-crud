package repository

import (
	"log"

	"github.com/PawelK2012/go-crud/database"
)

type Repository struct {
	Postgress database.ClientInterface
}

func New() (*Repository, error) {
	postgress, err := database.NewPostgressClient()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &Repository{Postgress: postgress}, err
}
