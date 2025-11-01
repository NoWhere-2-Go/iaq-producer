package models

import "time"

type SensorData struct {
	DeviceID    string    `json:"device_id"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	CO2         float64   `json:"co2"`
	VOC         float64   `json:"voc"`
	Timestamp   time.Time `json:"timestamp"`
}
