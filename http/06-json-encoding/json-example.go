package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var input user
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	if input.Name == "" || input.Age <= 0 {
		http.Error(w, "name and positive age are required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(input)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", createUser)

	fmt.Println("POST JSON to http://localhost:8080/users")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("server error:", err)
	}
}
