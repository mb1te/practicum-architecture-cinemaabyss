package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mb1te/practicum-architecture-cinemaabyss/microservices/events/internal"
	api "github.com/mb1te/practicum-architecture-cinemaabyss/microservices/events/internal/http"
	"github.com/mb1te/practicum-architecture-cinemaabyss/microservices/events/internal/http/handlers"
	"github.com/mb1te/practicum-architecture-cinemaabyss/microservices/events/internal/kafka"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := internal.NewConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	movieProducer := kafka.NewProducer(cfg.KafkaBrokers, "movie-events")
	defer movieProducer.Close()
	userProducer := kafka.NewProducer(cfg.KafkaBrokers, "user-events")
	defer userProducer.Close()
	paymentProducer := kafka.NewProducer(cfg.KafkaBrokers, "payment-events")
	defer paymentProducer.Close()

	movieConsumer := kafka.NewConsumer(cfg.KafkaBrokers, "movie-events")
	go movieConsumer.Run(ctx)
	userConsumer := kafka.NewConsumer(cfg.KafkaBrokers, "user-events")
	go userConsumer.Run(ctx)
	paymentConsumer := kafka.NewConsumer(cfg.KafkaBrokers, "payment-events")
	go paymentConsumer.Run(ctx)

	movieEventHandler := handlers.NewMovieEventHandler(movieProducer)
	userEventHandler := handlers.NewUserEventHandler(userProducer)
	paymentEventHandler := handlers.NewPaymentEventHandler(paymentProducer)

	router := api.NewRouter(movieEventHandler, userEventHandler, paymentEventHandler)
	endpoint := fmt.Sprintf(":%s", cfg.Port)
	if err := http.ListenAndServe(endpoint, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
