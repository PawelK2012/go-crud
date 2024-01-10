package main

import (
	"log"
	"os"

	"github.com/PawelK2012/go-crud/repository"
)

func main() {
	repository, err := repository.New()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Printf("%+v\n", repository)

	server := NewAPIServer(":3000", repository)
	server.Run()
}
