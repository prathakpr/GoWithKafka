package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func produceMessage(topic, message string) {
	writer := kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("key"),
			Value: []byte(message),
		},
	)
	if err != nil {
		log.Fatal("Failed to write message:", err)
	}
	fmt.Println("Message sent to Kafka:", message)
	writer.Close()
}

func main() {
	produceMessage("test-topic", "Document Created")
}
