package main

import (
	"fmt"
	"sync"
)

func launchGoroutines() {
	// Define a channel for communication
	channel := make(chan int)

	// Use a WaitGroup to ensure all goroutines complete
	var wg sync.WaitGroup

	// Increment the WaitGroup counter for the two writer goroutines
	wg.Add(2)

	// First writer goroutine
	go func() {
		defer wg.Done() // Decrement the counter when this goroutine completes
		for i := 1; i <= 10; i++ {
			channel <- i
		}
	}()

	// Second writer goroutine
	go func() {
		defer wg.Done() // Decrement the counter when this goroutine completes
		for i := 11; i <= 20; i++ {
			channel <- i
		}
	}()

	// Start a goroutine to close the channel once all writers are done
	go func() {
		wg.Wait()      // Wait for the writer goroutines to finish
		close(channel) // Close the channel to signal the reader
	}()

	// Reader goroutine
	for value := range channel {
		fmt.Println(value)
	}
}

func main() {
	launchGoroutines()
}
