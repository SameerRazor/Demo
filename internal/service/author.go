package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"Demo/internal/models"
)

func CreateAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var author models.Author
		err := c.ShouldBindJSON(&author)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		result := db.Create(&author)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a new author"})
			return
		}

		c.JSON(http.StatusCreated, author)
	}
}

func GetAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var authors []models.Author
		result := db.Find(&authors)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
			return
		}
		c.JSON(http.StatusOK, authors)
	}
}

func UpdateAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Author ID"})
			return
		}

		var author models.Author
		result := db.First(&author, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
			return
		}

		err = c.ShouldBindJSON(&author)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		result = db.Save(&author)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the author"})
			return
		}

		c.JSON(http.StatusOK, author)
	}
}

func GetAuthorParams(db *gorm.DB, params string) gin.HandlerFunc {
	return func(c *gin.Context) {
		param, err := strconv.Atoi(c.Param(params))
		if err != nil {

			param := c.Param(params)

			var authors []models.Author
			result := db.Find(&authors, param)
			if result.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
				return
			}
			c.JSON(http.StatusOK, authors)
			return
		}

		var authors []models.Author
		result := db.First(&authors, param)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
			return
		}
		c.JSON(http.StatusOK, authors)
	}
}

func DeleteAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
			return
		}

		var author models.Author
		result := db.First(&author, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
			return
		}

		result = db.Delete(&author)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the author"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Author deleted"})
	}
}
