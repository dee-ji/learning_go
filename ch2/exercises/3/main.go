/*
Write a program with three variables, one named b of type byte, one named smallI of type int32, and one named bigI of type uint64. Assign each variable the maximum legal value of its type; then add 1 to each variable. Print out their values.
*/
package main

import "fmt"

func main() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	// Before
	fmt.Println("Initial values")
	fmt.Printf("b (byte): %b\n", b)
	fmt.Printf("smallI (int32): %d\n", smallI)
	fmt.Printf("bigI (uint64): %d\n", bigI)
	// After adding 1, representing overflow affect
	b += 1
	smallI += 1
	bigI += 1
	fmt.Println("Adding 1 to each variable")
	fmt.Printf("b + 1 (byte): %b\n", b)
	fmt.Printf("smallI + 1 (int32): %d\n", smallI)
	fmt.Printf("bigI + 1 (uint64): %d\n", bigI)
}
