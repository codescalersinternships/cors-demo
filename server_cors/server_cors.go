package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Hello, World!"}`)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", helloHandler)

	// Enable CORS
	handler := cors.Default().Handler(mux)

	fmt.Println("Server is running on http://localhost:3000")
	http.ListenAndServe(":3000", handler)
}
