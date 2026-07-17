package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// read request info
	fmt.Println("method", r.Method)
	fmt.Println("Path:", r.URL.Path)

	//send response
	fmt.Fprint(w, "hello world")
}

func main() {

	// route "/" to handler
	http.HandleFunc("/", handler)

	fmt.Println("Server running on :8080")

	// starts the server
	http.ListenAndServe(":8080", nil)
}
