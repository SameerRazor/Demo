package router

import (
	"Demo/internal/author/service"
	"Demo/internal/book/service"
	"Demo/internal/genre/service"
	"Demo/internal/library/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	r.GET("/books", bookService.GetBooks(db))
	r.GET("/books/:params", bookService.GetBookParams(db, "params"))
	r.POST("/books", bookService.CreateBooks(db))
	r.PATCH("/books/:id", bookService.UpdateBooks(db))
	r.DELETE("/books/:id", bookService.DeleteBook(db))

	r.POST("/author", authorService.CreateAuthor(db))
	r.GET("/author", authorService.GetAuthor(db))
	r.GET("/getAuthor", authorService.GetAuthorParams(db))
	r.PATCH("/author/:id", authorService.UpdateAuthor(db))
	r.DELETE("/author/:id", authorService.DeleteAuthor(db))

	r.POST("/genre", genreService.CreateGenre(db))
	r.GET("/genre/:id", genreService.GetGenreById(db))
	r.PATCH("/genre/:id", genreService.UpdateGenre(db))
	r.DELETE("genre/:id", genreService.DeleteGenre(db))

	r.POST("/library", libraryService.StoreBook(db))
	r.GET("/library/:id", libraryService.GetPositionByID(db))
	r.GET("/author/:id", libraryService.GetBooksPositionByAuthor(db))
	r.DELETE("library/:id", libraryService.RemoveBook(db))

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
