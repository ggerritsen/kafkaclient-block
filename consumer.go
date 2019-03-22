package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// Record is the object stored in the kafka topic
type Record struct {
	name string
}

type consumer struct {
	r *kafka.Reader
}

// NewConsumer creates a consumer that consumes from a kafka topic
func NewConsumer(brokers []string, topic string) (*consumer, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(0)

	return &consumer{r}, nil
}

// Close closes the underlying kafka reader
func (c *consumer) Close() {
	c.r.Close()
}

// Consume consumes from the topic until an error is encountered
func (c *consumer) Consume() error {
	for {
		m, err := c.r.ReadMessage(context.Background())
		if err != nil {
			return err
		}
		fmt.Printf("Received message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	// return nil
}
