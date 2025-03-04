package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int) {
	for task := range tasks {
		fmt.Printf("Worker %d: Processing task %d\n", id, task)
		time.Sleep(500 * time.Millisecond) // Simulate slow work
		fmt.Printf("Worker %d: Finished task %d\n", id, task)
	}
}

func main() {
	// Create a buffered channel with capacity 3
	tasks := make(chan int, 3)

	// Start two workers
	go worker(1, tasks)
	go worker(2, tasks)

	// Start two workers
	for i := 1; i <= 5; i++ {
		fmt.Printf("manager: Assigning task %d\n", i)
		tasks <- i // Send task to buffered channel
		fmt.Printf("manager: Assigned task %d (buffer len: %d)\n", i, len(tasks))
		time.Sleep(100 * time.Millisecond) // Manager works faster than workers

	}

	// Close the channel after all tasks are assigned
	close(tasks)

	// Wait for workers to finish
	time.Sleep(3 * time.Second)
}
