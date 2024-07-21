package main

import (
	"log"
	"suram_backend/cmd/api"
)

func main() {
	server := api.NewApiServer(":8080", nil)
	if err := server.Start(); err != nil {
		log.Fatal()
	}
}
