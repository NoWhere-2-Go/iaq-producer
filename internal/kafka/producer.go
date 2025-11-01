package kafka

import (
	"context"
	"encoding/json"
	"iaq-producer/internal/models"
	"log"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
	topic  string
}

func NewProducer(brokers, topic string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
		topic: topic,
	}
}

func (p *Producer) Send(ctx context.Context, msg models.SensorData) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(msg.DeviceID),
		Value: data,
	})
}

func (p *Producer) Close() {
	if err := p.writer.Close(); err != nil {
		log.Printf("Failed to close producer: %v", err)
	}
}
