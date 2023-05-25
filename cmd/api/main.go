package main

import (
	"log"

	"Demo/config"
	"Demo/internal/book/models"
	"Demo/internal/author/models"
	"Demo/internal/genre/models"
	"Demo/internal/library/models"
	"Demo/internal/book/router"
	"Demo/internal/author/router"
	"Demo/internal/genre/router"
	"Demo/internal/library/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.LoadConfig()

	r := gin.Default()

	err := db.AutoMigrate(&bookModels.Book{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
		return
	}
	err = db.AutoMigrate(&authorModels.Author{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
		return
	}
	err = db.AutoMigrate(&genreModels.Genre{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}

	err = db.AutoMigrate(&libraryModels.Library{})
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}

	authorRouter.SetupRouter(r, db)
	bookRouter.SetupRouter(r, db)
	genreRouter.SetupRouter(r, db)
	libraryRouter.SetupRouter(r, db)

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
