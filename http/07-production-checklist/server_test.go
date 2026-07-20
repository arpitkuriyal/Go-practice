package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name, body string
		wantStatus int
	}{
		{"valid request", `{"name":"Arpit","email":"arpit@example.com"}`, http.StatusCreated},
		{"unknown field", `{"name":"Arpit","email":"arpit@example.com","role":"admin"}`, http.StatusBadRequest},
		{"missing field", `{"name":"Arpit"}`, http.StatusUnprocessableEntity},
		{"two JSON objects", `{"name":"Arpit","email":"arpit@example.com"} {}`, http.StatusBadRequest},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(test.body))
			response := httptest.NewRecorder()

			newServer().ServeHTTP(response, request)

			if got := response.Code; got != test.wantStatus {
				t.Fatalf("status = %d, want %d; body = %s", got, test.wantStatus, response.Body.String())
			}
			if contentType := response.Header().Get("Content-Type"); contentType != "application/json; charset=utf-8" {
				t.Fatalf("Content-Type = %q", contentType)
			}
		})
	}
}

func TestMethodAndHealthRoutes(t *testing.T) {
	server := newServer()

	wrongMethod := httptest.NewRecorder()
	server.ServeHTTP(wrongMethod, httptest.NewRequest(http.MethodGet, "/users", nil))
	if got := wrongMethod.Code; got != http.StatusMethodNotAllowed {
		t.Fatalf("GET /users status = %d, want %d", got, http.StatusMethodNotAllowed)
	}

	health := httptest.NewRecorder()
	server.ServeHTTP(health, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	if got := health.Code; got != http.StatusOK {
		t.Fatalf("GET /healthz status = %d, want %d", got, http.StatusOK)
	}
}
