package main

import (
	"iaq-producer/internal/config"
	"iaq-producer/internal/kafka"
	"iaq-producer/internal/mqtt"
)

func main() {
	cfg := config.Load()

	producer := kafka.NewProducer(cfg)
	defer producer.Close()

	subscriber := mqtt.NewSubscriber(cfg.MQTTBroker, cfg.MQTTTopic, producer)
	subscriber.Start()
}
