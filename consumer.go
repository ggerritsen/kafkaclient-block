package main

type consumer struct {
}

// Record is the object stored in the kafka topic
type Record struct {
	name string
}

// NewConsumer creates a consumer that consumes from a kafka topic
func NewConsumer() (*consumer, error) {
	return &consumer{}, nil
}
