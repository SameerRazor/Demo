package main

import (
	"Demo/config"
	"Demo/internal/entities/author"
	"Demo/internal/entities/book"
	"Demo/internal/entities/genre"
	"Demo/internal/entities/library"
	"Demo/router"
	"log"

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
	db.Migrator().CreateIndex(&book.Book{}, "idx_title_author_genre_pubdate")
	err = db.AutoMigrate(&author.Author{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
		return
	}
	db.Migrator().CreateIndex(&author.Author{}, "idx_authorname")
	err = db.AutoMigrate(&genre.Genre{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}
	db.Migrator().CreateIndex(&genre.Genre{}, "idx_genrename")
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
