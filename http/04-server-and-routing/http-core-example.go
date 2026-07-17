package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About Page")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)

	http.ListenAndServe(":8080", mux)
}
