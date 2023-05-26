package library

import (
	"net/http"
	"strconv"

	"Demo/internal/book"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StoreBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var library Library
		err := c.ShouldBindJSON(&library)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		var existingBook Library
		result := db.Where("aisle = ? AND level = ? AND position = ?", library.Aisle, library.Level, library.Position).First(&existingBook)
		if result.Error == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Position already occupied"})
			return
		}

		result = db.Create(&library)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store the book"})
			return
		}

		c.JSON(http.StatusCreated, library)
	}
}
func GetPositionByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookID := c.Param("id")

		var library Library
		if err := db.Where("id = ?", bookID).First(&library).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		position := library.Position

		c.JSON(http.StatusOK, gin.H{"position": position})
	}
}

func RemoveBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var library Library
		err := c.ShouldBindJSON(&library)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		result := db.Where("aisle = ? AND level = ? AND position = ?", library.Aisle, library.Level, library.Position).Delete(&library)
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found at the specified position"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book removed successfully"})
	}
}

func GetBooksPositionByAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
			return
		}

		var book []book.Book
		result := db.Table("books").Where("author_id = ?", authorID).Find(&book)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query books"})
			return
		}

		// Create a slice to store book positions
		var bookPositions []Library
		// c.JSON(http.StatusOK, len(book))

		// Retrieve book positions from the library based on book IDs
		for _, book := range book {
			var library Library
			result = db.Table("library").Where("id = ?", book.ID).First(&library)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve book position"})
				return
			}

			// Create a book position struct
			bookPos := Library{
				ID:       book.ID,
				Aisle:    library.Aisle,
				Level:    library.Level,
				Position: library.Position,
			}

			// Append book position to the slice
			bookPositions = append(bookPositions, bookPos)
		}

		c.JSON(http.StatusOK, bookPositions)
	}
}
