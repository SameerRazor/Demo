package lms

import (
	"Demo/internal/entities/book"
	"Demo/internal/entities/genre"
	errorcodes "Demo/internal/error"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetGenresByParams(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		offsetStr := c.DefaultQuery("offset", "0")
		offset, err := strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidOffsetValue})
			return
		}

		limit := 3
		params := c.Request.URL.Query()

		var genres []genre.Genre
		var result *gorm.DB
		if len(params) == 0 {
			result = db.Where("genres.is_deleted = ?", false).Limit(limit).Offset(offset).Find(&genres)
			if result.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": errorcodes.GenreNotFound})
				return
			}
		} else {
			for paramType, j := range params {
				switch paramType {
				case "author_id":
					for _, paramValue := range j {
						authorID, err := strconv.Atoi(paramValue)
						if err != nil {
							c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidAuthorId})
							return
						}

						result = db.Table("genres").
							Joins("INNER JOIN books ON genres.id = books.genre_id").
							Joins("INNER JOIN books ON authors.id = books.author_id").
							Where("authors.id = ?", authorID).Limit(limit).Offset(offset).
							Find(&genres)

					}
				// case "book_id":
				// 	for _, paramValue := range j {
				// 		bookID, err := strconv.Atoi(paramValue)
				// 		if err != nil {
				// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
				// 			return
				// 		}

				// 		result = db.Table("genres").
				// 		Joins("INNER JOIN books ON genres.id = books.genre_id").
				// 		Joins("INNER JOIN books ON authors.id = books.author_id").
				// 		Where("authors.id = ?", authorID).Limit(limit).Offset(offset).
				// 		Find(&genres)
				// 	}
				default:
					c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidParamType})
					return

				}

			}
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": errorcodes.GenreQueryFailure})
				return
			}
		}

		nextOffset := offset + limit

		c.JSON(http.StatusOK, gin.H{
			"genres":      genres,
			"next_offset": nextOffset,
		})
	}
}

func GetGenreById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidGenreId})
			return
		}

		var genre []genre.Genre
		result := db.Where("genres.is_deleted = ?", false).First(&genre, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": errorcodes.GenreNotFound})
			return
		}
		c.JSON(http.StatusOK, genre)
	}
}

func CreateGenre(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var genre genre.Genre
		err := c.ShouldBindJSON(&genre)
		if err != nil || genre.Genre == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidRequestPayload})
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": errorcodes.CreateNewGenreFailure})
			return
		}

		c.JSON(http.StatusCreated, genre)
	}
}

func UpdateGenre(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidGenreId})
			return
		}

		var genre genre.Genre
		result := db.First(&genre, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": errorcodes.GenreNotFound})
			return
		}

		err = c.ShouldBindJSON(&genre)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidRequestPayload})
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": errorcodes.FailedToUpdateGenre})
			return
		}

		c.JSON(http.StatusOK, genre)
	}
}

func DeleteGenre(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidGenreId})
			return
		}

		var genre genre.Genre
		result := db.First(&genre, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": errorcodes.GenreNotFound})
			return
		}

		genre.IsDeleted = true

		var boook book.Book
		result = db.Find(&boook, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": errorcodes.BookNotFound})
			return
		}

		boook.GenreId = 0

		db.Save(&boook)

		result = db.Delete(&genre)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errorcodes.GenreDeletionFailure})
			return
		}

		var books []book.Book
		result = db.Where("genre_id = ?", id).Find(&books)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errorcodes.FailedToFindBooksWithGenre})
			return
		}

		for i := range books {
			books[i].GenreId = 0
			result = db.Save(&books[i])
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": errorcodes.FailedToRemoveGenreAssociationFromBooks})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Genre deleted"})
	}
}
