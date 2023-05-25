package main

import (
	"log"

	"Demo/config"
	"Demo/internal/models"
	"Demo/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.LoadConfig()

	r := gin.Default()

	err := db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
		return
	}
	err = db.AutoMigrate(&models.Author{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
		return
	}
	err = db.AutoMigrate(&models.Genre{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}

	err = db.AutoMigrate(&models.Library{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}

	router.SetupRouter(r, db)

}
