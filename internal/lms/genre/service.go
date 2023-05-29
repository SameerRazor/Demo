package lms

import (
	"Demo/internal/entities/book"
	"Demo/internal/entities/genre"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetGenres(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var genres []genre.Genre
		result := db.Where("genres.is_deleted = ?", false).Find(&genres)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Books not found"})
			return
		}
		c.JSON(http.StatusOK, genres)
	}
}

func GetGenreById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
			return
		}

		var genre []genre.Genre
		result := db.Where("genres.is_deleted = ?", false).First(&genre, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
			return
		}
		c.JSON(http.StatusOK, genre)
	}
}

func CreateGenre(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var genre genre.Genre
		err := c.ShouldBindJSON(&genre)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		now := time.Now()
		genre.CreatedAt = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
			now.Year(),
			now.Month(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second())

		genre.UpdatedAt = genre.CreatedAt

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

		var genre genre.Genre
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

		now := time.Now()
		genre.UpdatedAt = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
			now.Year(),
			now.Month(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second())

		result = db.Updates(&genre)
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

		var genre genre.Genre
		result := db.First(&genre, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
			return
		}

		genre.IsDeleted = true

		var boook book.Book
		result = db.Find(&boook, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		boook.GenreId = 0

		db.Save(&boook)

		result = db.Delete(&genre)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the genre"})
			return
		}

		var books []book.Book
		result = db.Where("genre_id = ?", id).Find(&books)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find books with the genre"})
			return
		}

		for i := range books {
			books[i].GenreId = 0
			result = db.Save(&books[i])
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book genre association"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Genre deleted"})
	}
}
