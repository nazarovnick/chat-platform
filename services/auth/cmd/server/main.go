package main

import (
	"github.com/nazarovnick/chat-platform/services/auth/internal/controller/http"
	"log"
)

// main initializes the HTTP router and starts the server on specified port
//
// @title           Auth Service API
// @version         1.0
// @description     This is the authentication microservice API for the Chat Platform. It provides user registration, authentication, and related features.
// @host      		localhost:8080
// @BasePath  		/api/v1
// @securityDefinitions.basic  BasicAuth
func main() {
	app := http.NewRouter()
	addr := ":8080"
	log.Printf("Server running at %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
