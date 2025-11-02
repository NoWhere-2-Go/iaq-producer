package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	KafkaBrokers []string
	KafkaTopic   string
	ClientID     string
	MQTTBroker   string
	LogLevel     string
	Username     string
	Password     string
}

func Load() *Config {
	envFile, _ := godotenv.Read(".env")

	return &Config{
		KafkaBrokers: []string{envFile["KAFKA_BROKERS"]},
		KafkaTopic:   envFile["KAFKA_TOPIC"],
		ClientID:     envFile["CLIENT_ID"],
		MQTTBroker:   envFile["MQTT_BROKER"],
		LogLevel:     envFile["LOG_LEVEL"],
		Username:     envFile["KAFKA_USERNAME"],
		Password:     envFile["KAFKA_PASSWORD"],
	}
}
