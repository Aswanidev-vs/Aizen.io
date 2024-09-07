package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve static files (HTML, CSS) from the current directory
	http.Handle("/", http.FileServer(http.Dir(".")))

	// Start the server on port 8080
	log.Println("Starting server onhttp://127.0.0.1:3000/aizen/index.html")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
