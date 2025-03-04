package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel to pass coffee orders
	orders := make(chan string)

	// Cashier goroutine: send orders to the channel
	go func() {
		fmt.Println("Cashier: Taking order...")
		time.Sleep(1 * time.Second) // Simulate taking time
		orders <- "latte"           // Send order to channel
		fmt.Println("Cashier: Order sent!")
	}()

	// Barrista goroutine: receives orders from the channel
	go func() {
		fmt.Println("Barrista: Waiting for an order...")
		order := <-orders // Receive order from channel
		fmt.Println("Barrista: Making a", order)
	}()

	// Give the goroutines time to finish
	time.Sleep(2 * time.Second)
}
