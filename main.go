package main

import (
	"log"
	"net/http"
	"os"
)

const webRoot = "/var/www/html"

func main() {
	// Create web root if it doesn't exist
	if err := os.MkdirAll(webRoot, 0755); err != nil {
		log.Fatalf("Failed to create %s: %v", webRoot, err)
	}

	fs := http.FileServer(http.Dir(webRoot))

	log.Printf("Serving %s on :8080", webRoot)

	if err := http.ListenAndServe(":8080", fs); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
