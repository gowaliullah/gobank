package main

import (
	"fmt"
	"log"
)

func main() {

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Postgres store initialized: %+v\n", store)

	server := NewAPIServer(":8080", store)
	server.Run()
}
