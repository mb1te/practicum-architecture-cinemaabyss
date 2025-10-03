package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mb1te/practicum-architecture-cinemaabyss/microservices/proxy/internal/http/handlers"
)

func NewRouter(cfg handlers.TrafficSplitConfig) *gin.Engine {
	r := gin.Default()
	r.GET("/health", handlers.Health)
	r.Any("/api/*path", handlers.NewTrafficSplitHandler(cfg))
	return r
}
