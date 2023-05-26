package main

import (
	"Demo/config"
	"Demo/internal/book/models"
	"Demo/internal/book/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func MustJSONMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic("Failed to marshal JSON: " + err.Error())
	}
	return b
}

type CustomResponseWriter struct {
	*httptest.ResponseRecorder
}

func TestCreateBooks(t *testing.T) {

	db := config.LoadConfig()

	bookPayload := book.Book{
		ID: 1,
		Title: "The Merchant of Venice",
		AuthorName: "Sameer",
		AuthorId: 1,
		GenreName: "Romance",
		GenreId: 1,
		PublicationDate: "2054634354",
	}

	requestBody := gin.H{
		"id":          bookPayload.ID,
		"title":       bookPayload.Title,
		"author_name": bookPayload.AuthorName,
		"author_id":   bookPayload.AuthorId,
		"genre_name":  bookPayload.GenreName,
		"genre_id":    bookPayload.GenreId,
		"publication_date": bookPayload.PublicationDate,
	}
	requestJSON := map[string]interface{}{
		"json": requestBody,
	}

	requestBodyBytes := MustJSONMarshal(requestJSON)
	request, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")

	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	customResponseWriter := &CustomResponseWriter{response}

	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(customResponseWriter)
	context.Request = request

	bookService.CreateBooks(db)(context)

	assert.Equal(t, http.StatusCreated, response.Code)

	var createdBook book.Book
	result := db.Table("books").First(&createdBook)
	assert.NoError(t, result.Error)
	assert.Equal(t, bookPayload.AuthorName, createdBook.AuthorName)
	assert.Equal(t, bookPayload.GenreName, createdBook.GenreName)
}