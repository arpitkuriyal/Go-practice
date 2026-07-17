package authmiddleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	middleware := Middleware(map[string]string{"good-token": "arpit"})
	protected := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := UserFromContext(r.Context())
		if !ok {
			t.Fatal("authenticated user missing from context")
		}
		_ = json.NewEncoder(w).Encode(map[string]string{"user": user})
	}))

	tests := []struct {
		name, authorization string
		wantStatus          int
	}{
		{"missing token", "", http.StatusUnauthorized},
		{"invalid token", "Bearer wrong-token", http.StatusUnauthorized},
		{"valid token", "Bearer good-token", http.StatusOK},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/profile", nil)
			request.Header.Set("Authorization", test.authorization)
			response := httptest.NewRecorder()

			protected.ServeHTTP(response, request)
			if got := response.Code; got != test.wantStatus {
				t.Fatalf("status = %d, want %d; body = %s", got, test.wantStatus, response.Body.String())
			}
		})
	}
}
