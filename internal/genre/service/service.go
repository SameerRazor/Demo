package genreService

import (
	book "Demo/internal/book/models"
	"Demo/internal/genre/models"
	"net/http"
	"strconv"

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

		var genre []genre.Genre
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
		var genre genre.Genre
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

		var genre genre.Genre
		result := db.First(&genre, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
			return
		}

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
