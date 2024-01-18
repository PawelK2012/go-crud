package main

import (
	"log"
	"os"

	"github.com/PawelK2012/go-crud/repository/postgress"
)

func main() {
	repo, err := postgress.NewPostgress()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	server := NewAPIServer(":3000", repo)
	server.Run()
}
