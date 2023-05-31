package lms

import (
	"Demo/internal/entities/author"
	"Demo/internal/entities/book"
	"Demo/internal/entities/genre"
	"Demo/internal/lms"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCreateBooksSuccess(t *testing.T) {
	router, db := setup.InitializeTable()

	mockAuthor := author.Author{
		ID:          1,
		AuthorName:  "John Doe",
		Biography:   "Male",
		DateOfBirth: "2023-02-25",
		Nationality: "American",
		CreatedAt:   "2023-05-28 12:34:56",
		UpdatedAt:   "2023-05-28 12:34:56",
		IsDeleted:   false,
	}
	db.Create(&mockAuthor)

	mockGenre := genre.Genre{
		ID:        1,
		Genre:     "Horror",
		CreatedAt: "2023-05-28 12:34:56",
		UpdatedAt: "2023-05-28 12:34:56",
		IsDeleted: false,
	}
	db.Create(&mockGenre)

	router.POST("/books", CreateBooks(db))

	bookPayload := book.Book{
		Title:           "Merchant of Venice",
		AuthorName:      "John Doe",
		GenreName:       "Horror",
		PublicationDate: "1677312000",
	}

	payloadJSON, _ := json.Marshal(bookPayload)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(payloadJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, resp.Code)
	}

	var createdBook book.Book
	err := json.Unmarshal(resp.Body.Bytes(), &createdBook)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	assert.Equal(t, bookPayload.Title, createdBook.Title)
	assert.Equal(t, bookPayload.AuthorName, createdBook.AuthorName)
	assert.Equal(t, bookPayload.GenreName, createdBook.GenreName)
	assert.Equal(t, "2023-02-25", createdBook.PublicationDate)

	setup.DeleteTables(db)
}

func TestCreateBooksFail(t *testing.T) {
	router, db := setup.InitializeTable()

	mockAuthor := author.Author{
		ID:          1,
		AuthorName:  "John Doe",
		Biography:   "Male",
		DateOfBirth: "2023-02-25",
		Nationality: "American",
		CreatedAt:   "2023-05-28 12:34:56",
		UpdatedAt:   "2023-05-28 12:34:56",
		IsDeleted:   false,
	}
	db.Create(&mockAuthor)

	mockGenre := genre.Genre{
		ID:        1,
		Genre:     "Horror",
		CreatedAt: "2023-05-28 12:34:56",
		UpdatedAt: "2023-05-28 12:34:56",
		IsDeleted: false,
	}
	db.Create(&mockGenre)

	router.POST("/books", CreateBooks(db))

	bookPayload := book.Book{
		AuthorName:      "John Doe",
		GenreName:       "Horror",
		PublicationDate: "1677312000",
	}

	payloadJSON, _ := json.Marshal(bookPayload)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(payloadJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, resp.Code)
	}

	var createdBook book.Book
	err := json.Unmarshal(resp.Body.Bytes(), &createdBook)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	assert.Equal(t, bookPayload.Title, createdBook.Title)
	assert.Equal(t, bookPayload.AuthorName, createdBook.AuthorName)
	assert.Equal(t, bookPayload.GenreName, createdBook.GenreName)
	assert.Equal(t, "2023-02-25", createdBook.PublicationDate)

	setup.DeleteTables(db)
}

func TestGetBookByIdSuccess(t *testing.T) {
	router, db := setup.InitializeTable()
	router.GET("/books/:id", GetBookById(db))

	bookPayload := book.Book{
		ID:              1,
		Title:           "Merchant of Venice",
		AuthorName:      "John Doe",
		GenreName:       "Horror",
		PublicationDate: "2023-02-25",
		CreatedAt:       "2023-05-28 12:34:56",
		UpdatedAt:       "2023-05-28 12:34:56",
		IsDeleted:       false,
	}

	db.Create(bookPayload)

	req, _ := http.NewRequest("GET", "/books/1", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}

	var fetchedBooks []book.Book
	err := json.Unmarshal(resp.Body.Bytes(), &fetchedBooks)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	fetchedBook := fetchedBooks[0]
	assert.Equal(t, true, reflect.DeepEqual(bookPayload, fetchedBook))
	setup.DeleteTables(db)
}
func TestGetBookByIdFail(t *testing.T) {
	router, db := setup.InitializeTable()
	router.GET("/books/:id", GetBookById(db))

	bookPayload := book.Book{
		ID:              1,
		Title:           "Merchant of Venice",
		AuthorName:      "John Doe",
		GenreName:       "Horror",
		PublicationDate: "2023-02-25",
		CreatedAt:       "2023-05-28 12:34:56",
		UpdatedAt:       "2023-05-28 12:34:56",
		IsDeleted:       false,
	}

	db.Create(bookPayload)

	req, _ := http.NewRequest("GET", "/books/2", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}

	var fetchedBooks []book.Book
	err := json.Unmarshal(resp.Body.Bytes(), &fetchedBooks)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	fetchedBook := fetchedBooks[0]
	assert.Equal(t, true, reflect.DeepEqual(bookPayload, fetchedBook))
	setup.DeleteTables(db)
}