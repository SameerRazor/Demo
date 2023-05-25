package router

import (
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"Demo/internal/service"
)

func SetupRouter(r *gin.Engine, db *gorm.DB){
	r.GET("/books", service.GetBooks(db))
	r.GET("/books/:params", service.GetBookParams(db, "params"))
	r.POST("/books", service.CreateBooks(db))
	r.PATCH("/books/:id", service.UpdateBooks(db))
	r.DELETE("/books/:id", service.DeleteBook(db))

	r.POST("/author", service.CreateAuthor(db))
	r.GET("/author", service.GetAuthor(db))
	r.GET("/author/:params", service.GetAuthorParams(db, "params"))
	r.PATCH("/author/:id", service.UpdateAuthor(db))
	r.DELETE("/author/:id", service.DeleteAuthor(db))

	r.POST("/genre", service.CreateGenre(db))
	r.GET("/genre/:id", service.GetGenreById(db))
	r.PATCH("/genre/:id", service.UpdateGenre(db))
	r.DELETE("genre/:id", service.DeleteGenre(db))

	r.POST("/library", service.CreateLibrary(db))
	r.GET("/library/:id", service.GetLibrary(db))
	r.DELETE("library/:id", service.DeleteLibrary(db))

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
