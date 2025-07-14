package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

type Order struct {
	ID     string `json:"id"`
	Item   string `json:"item"`
	Amount int    `json:"amount"`
}

var inventory = map[string]int{
	"widget": 100,
	"gadget": 100,
}

func main() {
	log.Println("Starting inventory-service...")

	brokers := os.Getenv("KAFKA_BROKER")
	if brokers == "" {
		brokers = "kafka:9092"
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokers},
		Topic:    "orders",
		GroupID:  "inventory-service",
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
		updateInventory(order)
	}
}

func updateInventory(order Order) {
	if _, ok := inventory[order.Item]; !ok {
		log.Printf("Unknown item: %s", order.Item)
		return
	}
	if inventory[order.Item] < order.Amount {
		log.Printf("Not enough inventory for %s (requested: %d, available: %d)", order.Item, order.Amount, inventory[order.Item])
		return
	}
	inventory[order.Item] -= order.Amount
	log.Printf("Inventory updated for %s: %d remaining", order.Item, inventory[order.Item])
}
