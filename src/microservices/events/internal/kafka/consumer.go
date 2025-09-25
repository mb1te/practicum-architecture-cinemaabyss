package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic string) *Consumer {
	config := kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: fmt.Sprintf("%s-consumer", topic),
	}
	return &Consumer{reader: kafka.NewReader(config)}
}

func (c *Consumer) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("Error reading message: %+v", err)
			}
			log.Printf("Message consumed: %+v", string(msg.Value))
		}
	}
}
