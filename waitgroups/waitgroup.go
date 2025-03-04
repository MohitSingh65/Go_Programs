package main

import (
	"fmt"
	"sync"
	"time"
)

// processFile simulates processing a file and returns the line count

func processFile(filename string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when Done
	fmt.Printf("Processing %s...\n", filename)
	time.Sleep(1 * time.Second) // Simulate work
	fmt.Printf("Finished %s\n", filename)
}

func main() {
	files := []string{"file1.txt", "file2.txt", "file3.text"}
	var wg sync.WaitGroup

	// Start a goroutine for each file
	for _, file := range files {
		wg.Add(1) // Increment counter for each task
		go processFile(file, &wg)
	}

	fmt.Println("Main: Waiting for all files to be processed...")
	wg.Wait() // Block until all goroutines finish
	fmt.Print("Main: All files processed!")
}
