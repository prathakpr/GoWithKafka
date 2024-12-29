package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func consumeMessages(topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
		GroupID: "mongo-group",
	})

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Error while reading message:", err)
		}
		fmt.Printf("Received: %s\n", string(msg.Value))
	}
}

func main() {
	consumeMessages("test-topic")
}
