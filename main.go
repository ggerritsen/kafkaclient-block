package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Printf("Start of demonstration.\n")

	// TODO make new producer

	c, err := NewConsumer([]string{"localhost:9092"}, "test")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	go c.Consume()

	time.Sleep(15 * time.Second)

	fmt.Printf("End of demonstration.\n")
}
