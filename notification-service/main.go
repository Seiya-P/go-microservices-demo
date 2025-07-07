package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type Order struct {
	ID     string `json:"id"`
	Item   string `json:"item"`
	Amount int    `json:"amount"`
}

func main() {
	log.Println("Starting notification-service...")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "orders",
		GroupID:  "notification-service",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}
		var order Order
		if err := json.Unmarshal(m.Value, &order); err != nil {
			log.Printf("Invalid order event: %v", err)
			continue
		}
		log.Printf("Notification sent for order %s: %d x %s", order.ID, order.Amount, order.Item)
	}
}
