# 🌬️ IAQ Producer Service

A lightweight Go microservice that bridges **MQTT IoT sensor data** and a **Redpanda (Kafka-compatible) cluster**.  
It subscribes to air quality readings from an MQTT broker and publishes them to a Redpanda topic in real-time.

---

## 🚀 Features

- 🧩 **MQTT Subscriber** — listens to incoming sensor data
- 🧠 **JSON Deserialization** — validates and parses IAQ messages into Go structs
- 📤 **Kafka/Redpanda Producer** — forwards readings to the configured topic
- 🔒 **SASL/TLS Support** — optional authentication for secured clusters
- 🧰 **Modular Design** — clean internal package layout for extensibility
