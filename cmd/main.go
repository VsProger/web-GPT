package main

import (
	"Ex2_Week3/pkg/chatgpt"
	"Ex2_Week3/pkg/config"
	"Ex2_Week3/pkg/web"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.LoadConfig()

	if cfg.APIKey == "" {
		log.Fatal("API key must be set")
	}
	client := chatgpt.NewClient(cfg.APIKey, cfg.APIBaseURL)
	router := mux.NewRouter()

	// Setup routes with the client
	web.SetupRoutes(router, client)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	log.Println("Server starting on http://127.0.0.1:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
