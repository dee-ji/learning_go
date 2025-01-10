package main

import "fmt"

// Write a function called prefixer that has an input parameter of type string and returns a function that has an input
// parameter of type string and returns a string. The returned function should prefix its input with the string passed
// into prefixer.
func prefixer(prefix string) func(string) string {
	return func(s string) string {
		return prefix + " " + s
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
