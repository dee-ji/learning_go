package main

import (
	"fmt"
	"math/rand"
)

// Write a for loop that puts 100 random numbers between 0 and 100 int an int slice.
func main() {
	var randomNums []int
	for i := 0; i < 100; i++ {
		// using rand.Intn(101) will allow both 0 and 100 to be possible random numbers
		randomNums = append(randomNums, rand.Intn(101))
	}
	fmt.Printf("Random Numbers:\n%v\n", randomNums)
}
