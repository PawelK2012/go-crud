package main

import (
	"log"
	"os"

	"github.com/PawelK2012/go-crud/repository"
)

func main() {
	repo, err := repository.NewPostgress()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	server := NewAPIServer(":3000", repo)
	server.Run()
}
