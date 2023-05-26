package router

import (
	"Demo/internal/author"
	"Demo/internal/book"
	"Demo/internal/genre"
	"Demo/internal/library"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	r.GET("/books", book.GetBooks(db))
	r.GET("/books/:params", book.GetBookParams(db, "params"))
	r.POST("/books", book.CreateBooks(db))
	r.PATCH("/books/:id", book.UpdateBooks(db))
	r.DELETE("/books/:id", book.DeleteBook(db))

	r.POST("/author", author.CreateAuthor(db))
	r.GET("/author", author.GetAuthor(db))
	r.GET("/getAuthor", author.GetAuthorParams(db))
	r.PATCH("/author/:id", author.UpdateAuthor(db))
	r.DELETE("/author/:id", author.DeleteAuthor(db))

	r.POST("/genre", genre.CreateGenre(db))
	r.GET("/genre/:id", genre.GetGenreById(db))
	r.PATCH("/genre/:id", genre.UpdateGenre(db))
	r.DELETE("genre/:id", genre.DeleteGenre(db))

	r.POST("/library", library.StoreBook(db))
	r.GET("/library/:id", library.GetPositionByID(db))
	r.GET("/author/:id", library.GetBooksPositionByAuthor(db))
	r.DELETE("library/:id", library.RemoveBook(db))

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
