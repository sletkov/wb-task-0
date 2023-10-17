package main

import (
	"log"
	"sletkov/backend/wb-task-0/internal/app/apiserver"

	_ "github.com/lib/pq"
)

func main() {
	config := apiserver.NewConfig()

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

}
