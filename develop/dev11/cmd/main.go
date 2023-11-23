package main

import (
	"dev11/internal/server"
	"log"
)

func main() {
	err := server.StartServer(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
