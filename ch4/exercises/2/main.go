package main

import (
	"fmt"
	"math/rand"
)

// Loop over the slice you created in exercise 1. For each value in the slice, apply the following rules:
// 1. If the value is divisible by 2, print "Two!"
// 2. If the value is divisible by 3, print "Three!"
// 3. If the value is divisible by 2 and 3, print "Six!". Don't print anything else.
// 4. Otherwise, print "Never mind".
func main() {
	var randomNums []int
	for i := 0; i < 100; i++ {
		// using rand.Intn(101) will allow both 0 and 100 to be possible random numbers
		randomNums = append(randomNums, rand.Intn(101))
	}
	for _, num := range randomNums {
		if num%2 == 0 && num%3 == 0 {
			fmt.Printf("%d is divisible by both 2 and 3, Six!\n", num)
		} else if num%2 == 0 {
			fmt.Printf("%d is divisible by 2, Two!\n", num)
		} else if num%3 == 0 {
			fmt.Printf("%d is divisible by 3, Three!\n", num)
		} else {
			fmt.Printf("%d is not divisibable by 2 or 3, Nevermind...\n", num)
		}
	}
}
