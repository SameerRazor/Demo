package lms

import (
	"Demo/internal/entities/author"
	"Demo/internal/entities/book"
	errorcodes "Demo/internal/error"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var author author.Author
		err := c.ShouldBindJSON(&author)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidRequestPayload})
			return
		}

		parsedInteger, err := strconv.ParseInt(author.DateOfBirth, 10, 64)
		if err != nil {
			fmt.Println("Error parsing integer:", err)
			return
		}

		now := time.Now()
		author.CreatedAt = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
			now.Year(),
			now.Month(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second())

		author.UpdatedAt = author.CreatedAt

		epochTimeSeconds := parsedInteger

		epochTime := time.Unix(epochTimeSeconds, 0)

		epochDateString := epochTime.Format("2006-01-02")
		author.DateOfBirth = epochDateString

		result := db.Create(&author)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errorcodes.CreateNewAuthorFailure})
			return
		}
		if author.AuthorName == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errorcodes.EmptyAuthorName})
			return
		}

		c.JSON(http.StatusCreated, author)
	}
}

func GetAuthorById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidAuthorId})
			return
		}

		var genre []author.Author
		result := db.First(&genre, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": errorcodes.AuthorNotFound})
			return
		}
		c.JSON(http.StatusOK, genre)
	}
}

func UpdateAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidBookId})
			return
		}

		var author author.Author
		result := db.First(&author, id)
		if result.Error != nil || author.IsDeleted {
			c.JSON(http.StatusNotFound, gin.H{"error": errorcodes.BookNotFound})
			return
		}

		now := time.Now()
		author.UpdatedAt = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
			now.Year(),
			now.Month(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second())

		err = c.ShouldBindJSON(&author)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidRequestPayload})
			return
		}

		result = db.Save(&author)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errorcodes.FailedToUpdateBook})
			return
		}

		c.JSON(http.StatusOK, author)
	}
}

func GetAuthorParams(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := c.Request.URL.Query()
		var authors []author.Author
		var result *gorm.DB
		if len(params) == 0 {
			result = db.Where("authors.is_deleted = ?", false).Find(&authors)
			if result.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": errorcodes.AuthorNotFound})
				return
			}
		} else {
			for paramType, j := range params {
				switch paramType {
				case "genre":
					for _, paramValue := range j {
						genreID, err := strconv.Atoi(paramValue)
						if err != nil {
							c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidGenreId})
							return
						}

						result = db.Table("authors").
							Joins("INNER JOIN books ON authors.id = books.author_id").
							Joins("INNER JOIN genres ON genres.id = books.genre_id").
							Where("genres.id = ?", genreID).
							Find(&authors)

					}
				case "nationality":
					for _, paramValue := range j {
						result = db.Where("nationality = ?", paramValue).Find(&authors)
					}
				case "name":
					for _, paramValue := range j {
						result = db.Where("author_name LIKE ?", "%"+paramValue+"%").Find(&authors)
					}
				default:
					c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidParamType})
					return

				}

			}
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": errorcodes.AuthorQueryFailure})
				return
			}
		}

		c.JSON(http.StatusOK, authors)

	}
}

func DeleteAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errorcodes.InvalidAuthorId})
			return
		}
		var booksToDelete []book.Book
		db.Where("author_id = ?", id).Find(&booksToDelete)

		var author author.Author

		result := db.Find(&author, id)
		if result.Error != nil || author.IsDeleted {
			c.JSON(http.StatusNotFound, gin.H{"error": errorcodes.AuthorNotFound})
			return
		}

		author.IsDeleted = true

		result = db.Save(&author)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errorcodes.AuthorDeletionFailure})
			return
		}

		result = db.Table("books").Where("author_id = ?", id).Update("is_deleted", true)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errorcodes.BookDeletionFailure})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "author deleted"})
	}
}
