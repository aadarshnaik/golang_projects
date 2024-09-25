package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aadarshnaik/golang_projects/LostandFound/authentication/models"
)

func TestLoginUser(t *testing.T) {
	// Test valid credentials
	t.Run("Valid Credentials", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/login", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		userData := &models.User{
			Username:     "User6",
			Passwordhash: "User6pass",
		}

		jsonData, err := json.Marshal(userData)
		if err != nil {
			t.Fatal(err)
		}

		req.Body = io.NopCloser(bytes.NewBuffer(jsonData))

		w := httptest.NewRecorder()

		LoginUser(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		if w.Header().Get("Authorization") == "" {
			t.Errorf("Expected Authorization header to be set")
		}

		if w.Body.String() != "" {
			t.Errorf("Expected response body to be empty")
		}
	})

	// Test invalid credentials
	t.Run("Invalid Credentials", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/login", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		userData := &models.User{
			Username:     "User1",
			Passwordhash: "wrongpassword",
		}

		jsonData, err := json.Marshal(userData)
		if err != nil {
			t.Fatal(err)
		}

		req.Body = io.NopCloser(bytes.NewBuffer(jsonData))

		w := httptest.NewRecorder()

		LoginUser(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
		}
	})

	// Test invalid content type
	t.Run("Invalid Content Type", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/login", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "text/plain")

		w := httptest.NewRecorder()

		LoginUser(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}
