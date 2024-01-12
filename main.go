package main

import (
	"log"
	"os"

	"github.com/PawelK2012/go-crud/database"
)

func main() {
	db, err := database.NewPostgressClient()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Printf("%+v\n", db)

	server := NewAPIServer(":3000", db)
	server.Run()
}
