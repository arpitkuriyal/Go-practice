package main

import (
	"fmt"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	// /greet?name=Arpit
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	// Headers must be set before the body is written.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greeting)

	fmt.Println("Open http://localhost:8080/greet?name=Arpit")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("server error:", err)
	}
}
