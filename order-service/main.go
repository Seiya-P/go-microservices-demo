package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/segmentio/kafka-go"
)

type Order struct {
	ID     string `json:"id"`
	Item   string `json:"item"`
	Amount int    `json:"amount"`
}

func main() {
	log.Println("Starting order-service...")

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "orders",
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var order Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			log.Printf("Invalid order payload: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if order.ID == "" || order.Item == "" || order.Amount <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing or invalid order fields"))
			return
		}

		orderBytes, err := json.Marshal(order)
		if err != nil {
			log.Printf("Failed to marshal order: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		msg := kafka.Message{
			Key:   []byte(order.ID),
			Value: orderBytes,
			Time:  time.Now(),
		}
		if err := writer.WriteMessages(context.Background(), msg); err != nil {
			log.Printf("Failed to publish order to Kafka: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("Order published: %+v", order)
		w.WriteHeader(http.StatusCreated)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
