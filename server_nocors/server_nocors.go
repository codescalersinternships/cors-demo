package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Hello, World!"}`)
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	fmt.Println("Server is running on http://localhost:7000")
	http.ListenAndServe(":7000", nil)
}
