package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mb1te/practicum-architecture-cinemaabyss/microservices/events/internal/http/handlers"
)

func NewRouter(
	movieEventHandler *handlers.MovieEventHandler,
	userEventHandler *handlers.UserEventHandler,
	paymentEventHandler *handlers.PaymentEventHandler,
) *gin.Engine {
	r := gin.New()

	r.GET("/api/events/health", handlers.Health)
	r.POST("/api/events/movie", movieEventHandler.Create)
	r.POST("/api/events/user", userEventHandler.Create)
	r.POST("/api/events/payment", paymentEventHandler.Create)

	return r
}
