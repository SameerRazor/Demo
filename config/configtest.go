package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"Demo/internal/entities/book"
	"Demo/internal/entities/author"
	"Demo/internal/entities/genre"
	"Demo/internal/entities/library"
)

func LoadConfigTest() *gorm.DB {
	dsn := "root:J4C7ukpk@tcp(localhost:3306)/testapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	err = db.AutoMigrate(&book.Book{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}
	err = db.AutoMigrate(&author.Author{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}
	err = db.AutoMigrate(&genre.Genre{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}
	err = db.AutoMigrate(&library.Library{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}
	return db
}
