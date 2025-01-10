package main

import "fmt"

// Write a program that defines a variable named `greetings` of type slice of strings with the following values:
// "Hello", "Hola", "What's up?", "Que pasa?", "Howdy?". Create a subslice containing the first two values, a second
// subslice with the second, third and forth values; and a third subslice with the forth and fifth values. Print out
// all four slices.
func main() {
	var greetings = []string{"Hello", "Hola", "What's up?", "Que pasa?", "Howdy"}
	var subsliceOne = greetings[:2]   // first two values
	var subsliceTwo = greetings[1:4]  // second, third and forth values
	var subsliceThree = greetings[3:] // forth and fifth values

	// Print out all subslices
	fmt.Printf("%q\n", greetings)
	fmt.Println(subsliceOne, subsliceTwo, subsliceThree)
}
