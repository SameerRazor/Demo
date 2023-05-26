package bookService

import (
	"net/http"
	"strconv"

	"Demo/internal/author/models"
	"Demo/internal/genre/models"
	"Demo/internal/book/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var books []book.Book
		result := db.Find(&books)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Books not found"})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}

func GetBookParams(db *gorm.DB, params string) gin.HandlerFunc {
	return func(c *gin.Context) {
		param, err := strconv.Atoi(c.Param(params))
		if err != nil {

			param := c.Param(params)

			var books []book.Book
			result := db.Find(&books, param)
			if result.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
				return
			}
			c.JSON(http.StatusOK, books)
			return
		}

		var books []book.Book
		result := db.First(&books, param)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}

func CreateBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var book book.Book
		err := c.ShouldBindJSON(&book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		genrename := book.GenreName
		authorname := book.AuthorName

		var author author.Author
		result := db.Table("authors").
			Where("author_name = ?", authorname).
			First(&author)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve author"})
			return
		}

		book.AuthorId = author.ID

		var genre genre.Genre
		result = db.Table("genres").
			Where("genre = ?", genrename).
			First(&genre)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve genre"})
			return
		}

		book.GenreId = genre.ID

		result = db.Create(&book)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a new book"})
			return
		}

		c.JSON(http.StatusCreated, book)
	}
}

func UpdateBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}

		var book book.Book
		result := db.First(&book, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		err = c.ShouldBindJSON(&book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		result = db.Save(&book)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the book"})
			return
		}

		c.JSON(http.StatusOK, book)
	}
}

func DeleteBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}

		var book book.Book
		result := db.First(&book, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		result = db.Delete(&book)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the book"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
	}
}
