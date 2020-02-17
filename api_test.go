package main

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupMockedRouter(t *testing.T) (*gin.Engine, *sql.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db, err := gorm.Open("postgres", mockDB)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	router := setupRouter(db)

	return router, mockDB, mock
}

func TestHealth(t *testing.T) {
	router, db, _ := setupMockedRouter(t)
	defer db.Close()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	expectedCode := 200
	expectedBody := `{"status":"ok"}`
	responseCode := w.Code
	responseBody := w.Body.String()

	if responseCode != expectedCode {
		t.Errorf("Response code, expected %d but got %d", expectedCode, w.Code)
	}

	if !strings.Contains(responseBody, expectedBody) {
		t.Errorf("Response body, expected %s but got %s", expectedBody, responseBody)
	}
}

func TestGetDummy(t *testing.T) {
	router, db, mock := setupMockedRouter(t)
	defer db.Close()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/dummy", nil)
	router.ServeHTTP(w, req)

	expectedCode := 200
	// expectedBody := `{"status":"ok"}`
	responseCode := w.Code
	// responseBody := w.Body.String()

	if responseCode != expectedCode {
		t.Errorf("Response code, expected %d but got %d", expectedCode, w.Code)
	}

	fmt.Printf("%T, %s\n", mock, w.Body.String())
}
