package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	router := setupRouter()

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
