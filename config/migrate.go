package config

import (
	"Demo/internal/entities/author"
	"Demo/internal/entities/book"
	"Demo/internal/entities/genre"
	"Demo/internal/entities/library"
	"log"

	"gorm.io/gorm"
)

func AutoMigrateTable(db *gorm.DB) {
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
}
