// Write a program that declares var int i with value 20.
// Assign i to a float f. Print out both i and f.
package main

import "fmt"

func main() {
	var i int = 20
	fmt.Println(i)
	var f float32
	f = float32(i)
	fmt.Println(f)
}
