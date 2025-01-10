package main

import "fmt"

// Start a new program. In main, declare an int variable called total. Write a for loop that uses a variable named i to
// iterate from 0 (inclusive) to 10 (exclusive). The body of the for loop should be as follows:
// total := total + i
// fmt.Println(total)
// After the for loop, print out the value of total. What is the printed total? What is the likely bug in this code?
func main() {
	var total int
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
	fmt.Printf("The total is: %d", total) // The total will be 0
	// This will always by 0 because := only reuses variables that are declared in the current block
	// When using := be sure that you don't have any variables from an outer scope on the left hand side
	// unless you intend to shadow them.
}
