package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()

	WriteError(
		w,
		http.StatusNotFound,
		"Threat not found",
		"No threat with the specified MD5 was found",
	)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected application/json, got %s", contentType)
	}

	var response ErrorResponse

	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Invalid JSON response: %v", err)
	}

	if response.Message != "Threat not found" {
		t.Errorf("Unexpected message: %s", response.Message)
	}

	if response.Details != "No threat with the specified MD5 was found" {
		t.Errorf("Unexpected details: %s", response.Details)
	}
}
