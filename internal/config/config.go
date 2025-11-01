package config

import "os"

type Config struct {
	KafkaBrokers string
	KafkaTopic   string
	ClientID     string
	MQTTBroker   string
	LogLevel     string
}

func Load() *Config {
	return &Config{
		KafkaBrokers: os.Getenv("KAFKA_BROKERS"),
		KafkaTopic:   os.Getenv("KAFKA_TOPIC"),
		ClientID:     os.Getenv("CLIENT_ID"),
		MQTTBroker:   os.Getenv("MQTT_BROKER"),
		LogLevel:     os.Getenv("LOG_LEVEL"),
	}
}
