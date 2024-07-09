package main

import (
	"go-todo/cmd/storage"
	"log"
)

func main() {
	store, err := storage.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.InitDB(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer("8080", store)
	server.Run()
}
