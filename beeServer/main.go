package main

import (
	"bee/handlers"
	"bee/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := models.ConnectDB("user=vetrof dbname=beeDb host=localhost port=5433 sslmode=disable password=159753")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.POST("/register", handlers.CreateUserHandler(db))
	r.POST("/login", handlers.LoginHandler(db))

	r.Run(":8080")
}
