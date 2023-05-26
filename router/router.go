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

	r.POST("/authors", authorService.CreateAuthor(db))
	r.GET("/authors", authorService.GetAuthor(db))
	r.GET("/authors/:id", authorService.GetAuthorById(db))
	r.GET("/getAuthors", authorService.GetAuthorParams(db))
	r.PATCH("/authors/:id", authorService.UpdateAuthor(db))
	r.DELETE("/authors/:id", authorService.DeleteAuthor(db))

	r.POST("/genres", genreService.CreateGenre(db))
	r.GET("/genres/:id", genreService.GetGenreById(db))
	r.PATCH("/genres/:id", genreService.UpdateGenre(db))
	r.DELETE("genres/:id", genreService.DeleteGenre(db))

	r.POST("/libraries", libraryService.StoreBook(db))
	r.GET("/libraries/:id", libraryService.GetPositionByID(db))
	r.GET("/author/:id", libraryService.GetBooksPositionByAuthor(db))
	r.DELETE("libraries/:id", libraryService.RemoveBook(db))

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
