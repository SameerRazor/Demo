package lms

import (
	"Demo/internal/entities/author"
	"Demo/internal/lms"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAuthorSuccess(t *testing.T) {
	router, db := setup.InitializeTable()

	router.POST("/authors", CreateAuthor(db))

	authorPayload := author.Author{
		AuthorName:  "John Doe",
		Biography:   "Male",
		DateOfBirth: "1677312000",
		Nationality: "American",
	}

	payloadJSON, _ := json.Marshal(authorPayload)

	req, _ := http.NewRequest("POST", "/authors", bytes.NewBuffer(payloadJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, resp.Code)
	}

	var createdAuthor author.Author
	err := json.Unmarshal(resp.Body.Bytes(), &createdAuthor)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	assert.Equal(t, authorPayload.AuthorName, createdAuthor.AuthorName)
	assert.Equal(t, "2023-02-25", createdAuthor.DateOfBirth)
	assert.Equal(t, authorPayload.Nationality, createdAuthor.Nationality)
	assert.Equal(t, authorPayload.Biography, createdAuthor.Biography)

	setup.DeleteTables(db)
}

func TestCreateAuthorFail(t *testing.T) {
	router, db := setup.InitializeTable()

	router.POST("/authors", CreateAuthor(db))

	authorPayload := author.Author{
		Biography:   "Male",
		DateOfBirth: "1677312000",
		Nationality: "American",
	}

	payloadJSON, _ := json.Marshal(authorPayload)

	req, _ := http.NewRequest("POST", "/authors", bytes.NewBuffer(payloadJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, resp.Code)
	}

	var createdAuthor author.Author
	err := json.Unmarshal(resp.Body.Bytes(), &createdAuthor)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	assert.Equal(t, authorPayload.AuthorName, createdAuthor.AuthorName)
	assert.Equal(t, "2023-02-25", createdAuthor.DateOfBirth)
	assert.Equal(t, authorPayload.Nationality, createdAuthor.Nationality)
	assert.Equal(t, authorPayload.Biography, createdAuthor.Biography)

	setup.DeleteTables(db)
}

func TestGetAuthorByIdSuccess(t *testing.T) {
	router, db := setup.InitializeTable()

	router.GET("/authors/:id", GetAuthorById(db))

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

	req, _ := http.NewRequest("GET", "/authors/1", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}

	var fetchedAuthors []author.Author
	err := json.Unmarshal(resp.Body.Bytes(), &fetchedAuthors)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	fetchedAuthor := fetchedAuthors[0]
	assert.Equal(t, true, reflect.DeepEqual(mockAuthor, fetchedAuthor))

	setup.DeleteTables(db)
}
func TestGetAuthorByIdNotFound(t *testing.T) {
	router, db := setup.InitializeTable()

	router.GET("/authors/:id", GetAuthorById(db))

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

	req, _ := http.NewRequest("GET", "/authors/2", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}

	var fetchedAuthors []author.Author
	err := json.Unmarshal(resp.Body.Bytes(), &fetchedAuthors)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	fetchedAuthor := fetchedAuthors[0]
	assert.Equal(t, true, reflect.DeepEqual(mockAuthor, fetchedAuthor))

	setup.DeleteTables(db)
}
