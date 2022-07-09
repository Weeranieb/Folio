package main

import (
	"Folio/api"
	"log"
)

func main() {
	r := api.NewServer()

	err := r.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
