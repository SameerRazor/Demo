package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetBooks(t *testing.T) {
	router := gin.Default()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
	expectedBody := `[{"id":1,"title":"Book 1","author_name":"Author 1","author_id":1,"genre_name":"Genre 1","genre_id":1,"publicationdate":2023}]`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s but got %s", expectedBody, w.Body.String())
	}
}