package main

import (
	"gormex/db"
	"gormex/router"
	"log"
	"net/http"
	"os"
)

func main() {
	db.InitPostgres()

	// Server init
	r := router.Router()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}

	log.Println("Starting server at port", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
