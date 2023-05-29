package main

import (
	"log"
	"Demo/config"
	"Demo/router"
	"Demo/internal/entities/book"
	"Demo/internal/entities/author"
	"Demo/internal/entities/genre"
	"Demo/internal/entities/library"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.LoadConfig()

	r := gin.Default()

	err := db.AutoMigrate(&book.Book{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
		return
	}
	err = db.AutoMigrate(&author.Author{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
		return
	}
	err = db.AutoMigrate(&genre.Genre{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}

	err = db.AutoMigrate(&library.Library{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}

	router.SetupRouter(r, db)

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
