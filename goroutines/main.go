package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the counter when the goroutine completes
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	const numWorkers = 5

	wg.Add(numWorkers) // Set the counter to the number of goroutines we will spawn
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All workers have completed.")
}
