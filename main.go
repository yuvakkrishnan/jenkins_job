package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			log.Printf("Error getting hostname: %v", err)
			hostname = "unknown"
		}
		fmt.Fprintf(w, "Hello from GoLang service! Running on host: %s\n", hostname)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default port if not specified
	}

	log.Printf("GoLang service starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
