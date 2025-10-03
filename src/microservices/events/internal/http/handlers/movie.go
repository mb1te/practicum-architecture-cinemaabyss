package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mb1te/practicum-architecture-cinemaabyss/microservices/events/internal/kafka"
)

type MovieEvent struct {
	MovieID uint64 `json:"movie_id"`
	Title   string `json:"title"`
	Action  string `json:"action"`
	UserID  uint64 `json:"user_id"`
}

type MovieEventHandler struct {
	Producer *kafka.Producer
}

func NewMovieEventHandler(producer *kafka.Producer) *MovieEventHandler {
	return &MovieEventHandler{
		Producer: producer,
	}
}

func (h *MovieEventHandler) Create(c *gin.Context) {
	var payload MovieEvent
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payloadRaw, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.Producer.Produce(c.Request.Context(), payloadRaw); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}
