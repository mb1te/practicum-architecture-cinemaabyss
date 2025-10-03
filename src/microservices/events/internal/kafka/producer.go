package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	config := kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	}
	return &Producer{writer: kafka.NewWriter(config)}
}

func (p *Producer) Produce(ctx context.Context, message json.RawMessage) error {
	return p.writer.WriteMessages(ctx, kafka.Message{Value: message})
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
