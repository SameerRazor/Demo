package service

import (
	"net/http"
	"strconv"

	"Demo/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetGenreById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
			return
		}

		var genre []models.Genre
		result := db.First(&genre, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
			return
		}
		c.JSON(http.StatusOK, genre)
	}
}

func CreateGenre(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var genre models.Genre
		err := c.ShouldBindJSON(&genre)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		result := db.Create(&genre)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a new genre"})
			return
		}

		c.JSON(http.StatusCreated, genre)
	}
}

func UpdateGenre(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Genre ID"})
			return
		}

		var genre models.Genre
		result := db.First(&genre, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
			return
		}

		err = c.ShouldBindJSON(&genre)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		result = db.Save(&genre)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the genre"})
			return
		}

		c.JSON(http.StatusOK, genre)
	}
}

func DeleteGenre(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
			return
		}

		var genre models.Genre
		result := db.First(&genre, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
			return
		}

		result = db.Delete(&genre)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the Genre"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Genre deleted"})
	}
}
