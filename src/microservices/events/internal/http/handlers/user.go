package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mb1te/practicum-architecture-cinemaabyss/microservices/events/internal/kafka"
)

type UserEvent struct {
	UserID    uint64 `json:"user_id"`
	Username  string `json:"username"`
	Action    string `json:"action"`
	Timestamp string `json:"timestamp"`
}

type UserEventHandler struct {
	Producer *kafka.Producer
}

func NewUserEventHandler(producer *kafka.Producer) *UserEventHandler {
	return &UserEventHandler{
		Producer: producer,
	}
}

func (h *UserEventHandler) Create(c *gin.Context) {

	var payload UserEvent
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
