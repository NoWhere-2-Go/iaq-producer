package main

import (
	"iaq-producer/internal/app"
	"iaq-producer/internal/config"
	"iaq-producer/internal/kafka"
	"iaq-producer/internal/models"
	"log"
	"time"
)

func main() {
	cfg := config.Load()
	producer := kafka.NewProducer(cfg.KafkaBrokers, cfg.KafkaTopic)
	defer producer.Close()

	service := app.New(producer)

	data := models.SensorData{
		DeviceID:    "sensor-1",
		Temperature: 24.5,
		Humidity:    45.2,
		CO2:         410,
		VOC:         0.12,
		Timestamp:   time.Now(),
	}

	service.PublishReading(data)

	log.Println("✅ IAQ producer running...")
}
