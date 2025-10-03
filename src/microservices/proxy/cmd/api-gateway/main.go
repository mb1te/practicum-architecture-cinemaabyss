package main

import (
	"fmt"
	"log"
	"os"

	api "github.com/mb1te/practicum-architecture-cinemaabyss/microservices/proxy/internal/http"
	"github.com/mb1te/practicum-architecture-cinemaabyss/microservices/proxy/internal/http/handlers"
)

func main() {
	cfg, err := handlers.NewTrafficSplitConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	r := api.NewRouter(cfg)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
