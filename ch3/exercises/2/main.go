package main

import "fmt"

// Write a program that defines a string variable called `message` with the value "Hi 👨 and 🚀" and prints the
// forth rune in it as a character, not a number
func main() {
	var message string = "Hi 👨 and 👩"
	//fmt.Println(message)
	var forthRune = []rune(message)[3]
	fmt.Printf("The 4th rune in the message '%s' is %c\n", message, forthRune)
	fmt.Println(string(forthRune))
}
