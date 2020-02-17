package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	db, err := gorm.Open("postgres", mockDB)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	router := setupRouter(db)

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
