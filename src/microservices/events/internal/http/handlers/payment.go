package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mb1te/practicum-architecture-cinemaabyss/microservices/events/internal/kafka"
)

type PaymentEvent struct {
	PaymentID  uint64  `json:"payment_id"`
	UserID     uint64  `json:"user_id"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
	Timestamp  string  `json:"timestamp"`
	MethodType string  `json:"method_type"`
}

type PaymentEventHandler struct {
	Producer *kafka.Producer
}

func NewPaymentEventHandler(producer *kafka.Producer) *PaymentEventHandler {
	return &PaymentEventHandler{
		Producer: producer,
	}
}

func (h *PaymentEventHandler) Create(c *gin.Context) {

	var payload PaymentEvent
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
