package main

import (
	"fmt"
	"net/http"
)

// users handles more than one HTTP method for the same URL.
func users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "list users")
	case http.MethodPost:
		fmt.Fprintln(w, "create user")
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", users)

	fmt.Println("Try GET or POST http://localhost:8080/users")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("server error:", err)
	}
}
