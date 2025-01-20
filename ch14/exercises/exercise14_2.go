package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// randomAddWithTimeout runs the number generation simulation.
// It returns the sum, number of iterations, and the reason for ending (timeout or number reached).
func randomAddWithTimeout(timeout time.Duration) (int, int, string) {
	// Create a context with the specified timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	sum := 0
	iterations := 0
	reason := ""

	for {
		select {
		case <-ctx.Done(): // Check if the context's timeout or cancellation has occurred
			reason = "Timeout"
			return sum, iterations, reason
		default:
			// Generate a random number between 0 (inclusive) and 100,000,000 (exclusive)
			num := rand.Intn(100_000_000)
			sum += num
			iterations++

			// Check if the generated number is 1234
			if num == 1234 {
				reason = "Number 1234 reached"
				return sum, iterations, reason
			}
		}
	}
}

func main() {
	// Run the simulation with a 2-second timeout
	sum, iterations, reason := randomAddWithTimeout(2 * time.Second)

	// Print results
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Iterations: %d\n", iterations)
	fmt.Printf("Reason for ending: %s\n", reason)
}
