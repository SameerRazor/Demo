package lms

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"Demo/internal/entities/author"
	"Demo/internal/entities/book"
	"Demo/internal/entities/genre"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var books []book.Book
		result := db.Where("books.is_deleted = ?", false).Find(&books)
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

		now := time.Now()
		book.CreatedAt = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
			now.Year(),
			now.Month(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second())

		book.UpdatedAt = book.CreatedAt

		var author author.Author
		result := db.Table("authors").
			Where("author_name = ?", authorname).
			First(&author)
		if result.Error != nil || author.IsDeleted {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve author"})
			return
		}

		book.AuthorId = author.ID

		parsedInteger, err := strconv.ParseInt(book.PublicationDate, 10, 64)
		if err != nil {
			fmt.Println("Error parsing integer:", err)
			return
		}

		epochTimeSeconds := parsedInteger

		epochTime := time.Unix(epochTimeSeconds, 0)

		epochDateString := epochTime.Format("2006-01-02")
		book.PublicationDate = epochDateString

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

		now := time.Now()
		book.UpdatedAt = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
			now.Year(),
			now.Month(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second())

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
		if result.Error != nil || book.IsDeleted {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		book.IsDeleted = true
		result = db.Save(&book)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the book"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
	}
}
