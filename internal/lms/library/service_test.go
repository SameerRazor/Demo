package lms

import (
	"Demo/internal/entities/book"
	"Demo/internal/entities/library"
	"Demo/internal/lms"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreBookSuccess(t *testing.T) {
	router, db := setup.InitializeTable()

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

	router.POST("/libraries", StoreBook(db))

	mockBook := library.Library{
		Book_ID:  1,
		Aisle:    1,
		Level:    2,
		Position: 3,
	}

	payloadJSON, _ := json.Marshal(mockBook)

	req, _ := http.NewRequest("POST", "/libraries", bytes.NewBuffer(payloadJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, resp.Code)
	}

	var createdBook library.Library
	err := json.Unmarshal(resp.Body.Bytes(), &createdBook)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	assert.Equal(t, mockBook.Book_ID, createdBook.Book_ID)
	assert.Equal(t, mockBook.Aisle, createdBook.Aisle)
	assert.Equal(t, mockBook.Level, createdBook.Level)
	assert.Equal(t, mockBook.Position, createdBook.Position)

	setup.DeleteTables(db)
}
func TestStoreBookFail(t *testing.T) {
	router, db := setup.InitializeTable()

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

	router.POST("/libraries", StoreBook(db))

	mockBook := library.Library{
		Aisle:    1,
		Level:    2,
		Position: 3,
	}

	payloadJSON, _ := json.Marshal(mockBook)

	req, _ := http.NewRequest("POST", "/libraries", bytes.NewBuffer(payloadJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, resp.Code)
	}

	var createdBook library.Library
	err := json.Unmarshal(resp.Body.Bytes(), &createdBook)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	assert.Equal(t, mockBook.Book_ID, createdBook.Book_ID)
	assert.Equal(t, mockBook.Aisle, createdBook.Aisle)
	assert.Equal(t, mockBook.Level, createdBook.Level)
	assert.Equal(t, mockBook.Position, createdBook.Position)

	setup.DeleteTables(db)
}

func TestGetPositionByIDSuccess(t *testing.T) {
	router, db := setup.InitializeTable()
	router.GET("/libraries/:id", GetPositionByID(db))
	mockBook := library.Library{
		Book_ID:  1,
		Aisle:    1,
		Level:    2,
		Position: 3,
	}

	db.Create(&mockBook)

	req, _ := http.NewRequest("GET", "/libraries/1", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}

	var fetchedBook library.Library
	err := json.Unmarshal(resp.Body.Bytes(), &fetchedBook)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}
	assert.Equal(t, true, reflect.DeepEqual(mockBook, fetchedBook))

	setup.DeleteTables(db)
}
func TestGetPositionByIDFail(t *testing.T) {
	router, db := setup.InitializeTable()
	router.GET("/libraries/:id", GetPositionByID(db))

	mockBook := library.Library{
		Book_ID:  1,
		Aisle:    1,
		Level:    2,
		Position: 3,
	}

	db.Create(&mockBook)

	req, _ := http.NewRequest("GET", "/libraries/2", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}

	var fetchedBook library.Library
	err := json.Unmarshal(resp.Body.Bytes(), &fetchedBook)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}
	assert.Equal(t, true, reflect.DeepEqual(mockBook, fetchedBook))

	setup.DeleteTables(db)
}
