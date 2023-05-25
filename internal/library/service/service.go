package service

import (
	"net/http"
	"strconv"

	"Demo/internal/library/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateLibrary(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var library libraryModels.Library
		err := c.ShouldBindJSON(&library)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		result := db.Create(&library)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store a book in library"})
			return
		}

		c.JSON(http.StatusCreated, library)
	}
}
func GetLibrary(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}

		var position []libraryModels.Library
		result := db.First(&position, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusOK, position)
	}
}

func DeleteLibrary(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}

		var book libraryModels.Library
		result := db.First(&book, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		result = db.Delete(&book)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove the Book from it's position"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book removed from it's position"})
	}
}