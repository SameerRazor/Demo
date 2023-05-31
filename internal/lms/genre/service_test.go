package lms

import (
	"Demo/internal/entities/genre"
	"Demo/internal/lms"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGenreSuccess(t *testing.T) {
	router, db := setup.InitializeTable()

	router.POST("/genres", CreateGenre(db))

	genrePayload := genre.Genre{
		Genre: "Horror",
	}

	payloadJSON, _ := json.Marshal(genrePayload)

	req, _ := http.NewRequest("POST", "/genres", bytes.NewBuffer(payloadJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, resp.Code)
	}

	var createdGenre genre.Genre
	err := json.Unmarshal(resp.Body.Bytes(), &createdGenre)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	assert.Equal(t, genrePayload.Genre, createdGenre.Genre)

	setup.DeleteTables(db)
}

func TestCreateGenreFail(t *testing.T) {
	router, db := setup.InitializeTable()

	router.POST("/genres", CreateGenre(db))

	genrePayload := genre.Genre{
		Genre: "",
	}

	payloadJSON, _ := json.Marshal(genrePayload)

	req, _ := http.NewRequest("POST", "/genres", bytes.NewBuffer(payloadJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, resp.Code)
	}

	var createdGenre genre.Genre
	err := json.Unmarshal(resp.Body.Bytes(), &createdGenre)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	assert.Equal(t, genrePayload.Genre, createdGenre.Genre)

	setup.DeleteTables(db)
}

func TestGenreByIdSuccess(t *testing.T) {
	router, db := setup.InitializeTable()

	router.GET("/genres/:id", GetGenreById(db))

	mockGenre := genre.Genre{
		ID:    1,
		Genre: "Horror",
	}

	db.Create(&mockGenre)

	req, _ := http.NewRequest("GET", "/genres/1", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}

	var fetchedGenres []genre.Genre
	err := json.Unmarshal(resp.Body.Bytes(), &fetchedGenres)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	fetchedGenre := fetchedGenres[0]
	assert.Equal(t, true, reflect.DeepEqual(mockGenre, fetchedGenre))

	setup.DeleteTables(db)
}
func TestGenreByIdFail(t *testing.T) {
	router, db := setup.InitializeTable()

	router.GET("/genres/:id", GetGenreById(db))

	mockGenre := genre.Genre{
		ID:    1,
		Genre: "Horror",
	}

	db.Create(&mockGenre)

	req, _ := http.NewRequest("GET", "/genres/2", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}

	var fetchedGenres []genre.Genre
	err := json.Unmarshal(resp.Body.Bytes(), &fetchedGenres)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	fetchedGenre := fetchedGenres[0]
	assert.Equal(t, true, reflect.DeepEqual(mockGenre, fetchedGenre))

	setup.DeleteTables(db)
}