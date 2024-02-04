// main.go
package main

import (
	"net/http"
	"your-project/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Authentication routes
	r.HandleFunc("/signup", handlers.SignUpHandler).Methods("POST")
	// Add other authentication routes

	// Swipe route
	r.HandleFunc("/swipe", handlers.SwipeHandler).Methods("POST")

	// Start the server
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
