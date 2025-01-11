package main

import (
	"fmt"
	"time"
)

// Write a program that builds []Person with 10,000,000 entries (they could all be the same names and ages). See how
// long it takes to run. Change the value of GOGC and see how that affects the time it takes for the program to
// complete. Set the environment variable GODEBUG=gctrace=1 to see when garbage collections happen and see how changing
// GOGC changes the number of garbage collections.
// What happens if you create the slice with a capacity of 10,000,000?

type Person struct {
	name string
	age  int
}

func main() {
	start := time.Now()
	// Define a []Person without capacity
	var people []Person

	for i := 0; i < 10_000_000; i++ {
		people = append(people, Person{"DJ", 30})
	}
	duration := time.Since(start)
	fmt.Printf("Time taken without preallocated capacity: %v\n", duration)

	startWithCapacity := time.Now()
	// Make a []Person of 10,000,000 capacity
	peopleWithCapacity := make([]Person, 0, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		peopleWithCapacity = append(peopleWithCapacity, Person{"DJ", 30})
	}
	duration = time.Since(startWithCapacity)
	fmt.Printf("Time taken with preallocated capacity: %v\n", duration)
	// Results without defined capacity
	// $ time GOGC=10 GODEBUG=gctrace=1 go run exercise6_3.go
	// GOGC=10 GODEBUG=gctrace=1 go run exercise6_3.go  3.60s user 1.83s system 172% cpu 3.155 total
	// $ time GOGC=100 GODEBUG=gctrace=1 go run exercise6_3.go
	// GOGC=100 GODEBUG=gctrace=1 go run exercise6_3.go  2.58s user 1.45s system 155% cpu 2.595 total
	// $ time GOGC=1000 GODEBUG=gctrace=1 go run exercise6_3.go
	// GOGC=1000 GODEBUG=gctrace=1 go run exercise6_3.go  0.89s user 0.64s system 92% cpu 1.644 total

	// Results *with* defined capacity
	// $ time GOGC=10 GODEBUG=gctrace=1 go run exercise6_3.go
	// GOGC=10 GODEBUG=gctrace=1 go run exercise6_3.go  0.56s user 0.89s system 171% cpu 0.842 total
	// $ time GOGC=100 GODEBUG=gctrace=1 go run exercise6_3.go
	// GOGC=100 GODEBUG=gctrace=1 go run exercise6_3.go  0.45s user 0.86s system 163% cpu 0.800 total
	// $ time GOGC=1000 GODEBUG=gctrace=1 go run exercise6_3.go
	// GOGC=1000 GODEBUG=gctrace=1 go run exercise6_3.go  0.41s user 0.55s system 124% cpu 0.769 total
}
