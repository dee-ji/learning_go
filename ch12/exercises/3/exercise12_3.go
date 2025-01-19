package main

import (
	"fmt"
	"math"
	"sync"
)

var (
	once    sync.Once
	sqrtMap map[int]float64
)

func buildSqrtMap() {
	sqrtMap = make(map[int]float64, 100000)
	for i := 0; i < 100000; i++ {
		sqrtMap[i] = math.Sqrt(float64(i))
	}
}

func getSqrtMap() map[int]float64 {
	once.Do(buildSqrtMap) // Ensures the map is built only once
	return sqrtMap
}

func main() {
	// Use the cached map to look up square roots for every 1,000th number
	for i := 0; i < 100000; i += 1000 {
		sqrtValue := getSqrtMap()[i]
		fmt.Printf("Square root of %d: %.5f\n", i, sqrtValue)
	}
}
