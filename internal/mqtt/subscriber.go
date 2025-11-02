package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"iaq-producer/internal/kafka"
	"iaq-producer/internal/models"
)

type Subscriber struct {
	client   mqtt.Client
	producer *kafka.Producer
	topic    string
}

func NewSubscriber(broker, topic string, producer *kafka.Producer) *Subscriber {
	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetClientID("iaq-mqtt-bridge").
		SetCleanSession(true)

	sub := &Subscriber{
		producer: producer,
		topic:    topic,
	}
	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println("✅ Connected to MQTT broker")

		if token := c.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			var data models.SensorData
			if err := json.Unmarshal(msg.Payload(), &data); err != nil {
				log.Printf("⚠️  Invalid JSON received on %s: %v\n", msg.Topic(), err)
				return
			}

			fmt.Printf("MQTT message | Device: %s | Temp: %.2f°C | Humidity: %.2f%%\n",
				data.DeviceID, data.Temperature, data.Humidity)

			if err := sub.producer.Send(context.Background(), data); err != nil {
				log.Printf("❌ Kafka send failed: %v\n", err)
			} else {
				log.Printf("✅ Forwarded %s to Kafka\n", data.DeviceID)
			}
		}); token.Wait() && token.Error() != nil {
			log.Fatalf("❌ Failed to subscribe: %v", token.Error())
		}
	}

	sub.client = mqtt.NewClient(opts)
	return sub
}

func (s *Subscriber) Start() {
	if token := s.client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("❌ Failed to connect to MQTT broker: %v", token.Error())
	}
	fmt.Println("MQTT subscriber running...")
	select {}
}
