package main

import (
	"github.com/nazarovnick/chat-platform/services/auth/internal/controller/http"
	"log"
)

// main initializes the HTTP router and starts the server on specified port
func main() {
	app := http.NewRouter()
	addr := ":8080"
	log.Printf("Server running at %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
