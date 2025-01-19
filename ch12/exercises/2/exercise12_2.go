package main

import (
	"fmt"
	"sync"
)

func launchGoroutines() {
	// Create two channels for the two goroutines
	channel1 := make(chan int)
	channel2 := make(chan int)

	// Use a WaitGroup to synchronize the completion of the goroutines
	var wg sync.WaitGroup

	// Increment the WaitGroup counter for the two writer goroutines
	wg.Add(2)

	// First goroutine writes 10 numbers to channel1
	go func() {
		defer wg.Done() // Decrement the counter when this goroutine completes
		for i := 1; i <= 10; i++ {
			channel1 <- i
		}
		close(channel1) // Close the channel to signal no more values
	}()

	// Second goroutine writes 10 numbers to channel2
	go func() {
		defer wg.Done() // Decrement the counter when this goroutine completes
		for i := 1; i <= 10; i++ {
			channel2 <- i
		}
		close(channel2) // Close the channel to signal no more values
	}()

	// Start a goroutine to wait for writers to finish
	go func() {
		wg.Wait()       // Wait for both goroutines to complete
		close(channel1) // Close channel1 in case it's not already closed
		close(channel2) // Close channel2 in case it's not already closed
	}()

	// Use a for-select loop to read from both channels
	for {
		select {
		case val, ok := <-channel1:
			if ok {
				fmt.Printf("Goroutine 1 wrote: %d\n", val)
			} else {
				channel1 = nil // Set to nil to avoid further reads
			}
		case val, ok := <-channel2:
			if ok {
				fmt.Printf("Goroutine 2 wrote: %d\n", val)
			} else {
				channel2 = nil // Set to nil to avoid further reads
			}
		}

		// Break the loop when both channels are nil
		if channel1 == nil && channel2 == nil {
			break
		}
	}
}

func main() {
	launchGoroutines()
}
