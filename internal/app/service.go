package app

import (
	"context"
	"iaq-producer/internal/kafka"
	"iaq-producer/internal/models"
	"log"
	"time"
)

type Service struct {
	producer *kafka.Producer
}

func New(producer *kafka.Producer) *Service {
	return &Service{producer: producer}
}

func (s *Service) PublishReading(reading models.SensorData) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.producer.Send(ctx, reading); err != nil {
		log.Printf("Failed to publish reading: %v", err)
	} else {
		log.Printf("Published reading from %s", reading.DeviceID)
	}
}
