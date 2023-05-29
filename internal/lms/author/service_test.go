package lms

import (
	"Demo/config"
	"Demo/internal/entities/author"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateAuthor(t *testing.T) {
	router := gin.Default()

	db := config.LoadConfig()

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

	// assert.Equal(t, )
	if createdAuthor.AuthorName != authorPayload.AuthorName {
		t.Errorf("Expected AuthorName to be %s but got %s", authorPayload.AuthorName, createdAuthor.AuthorName)
	}
	if createdAuthor.DateOfBirth != "2023-02-25" {
		t.Errorf("Expected DateOfBirth to be %s but got %s", "2023-06-26", createdAuthor.DateOfBirth)
	}
	if createdAuthor.Nationality != authorPayload.Nationality {
		t.Errorf("Expected Nationality to be %s but got %s", authorPayload.Nationality, createdAuthor.Nationality)
	}
	if createdAuthor.Biography != authorPayload.Biography {
		t.Errorf("Expected Biography to be %s but got %s", authorPayload.Biography, createdAuthor.Biography)
	}
}
