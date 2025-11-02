package kafka

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"iaq-producer/internal/config"
	"iaq-producer/internal/models"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type Producer struct {
	writer *kafka.Writer
	topic  string
}

func NewProducer(config *config.Config) *Producer {
	mechanism, err := scram.Mechanism(scram.SHA256, config.Username, config.Password)
	if err != nil {
		log.Fatalf("Failed to create SASL mechanism: %v", err)
	}

	temp := &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(config.KafkaBrokers...),
			Topic:    config.KafkaTopic,
			Balancer: &kafka.LeastBytes{},
			Transport: &kafka.Transport{
				SASL: mechanism,
				TLS:  &tls.Config{},
			},
		},
		topic: config.KafkaTopic,
	}
	return temp
}

func (p *Producer) Send(ctx context.Context, msg models.SensorData) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(msg.DeviceID),
		Value: data,
		Time:  time.Now(),
	})
}

func (p *Producer) Close() {
	if err := p.writer.Close(); err != nil {
		log.Printf("Failed to close producer: %v", err)
	}
}
