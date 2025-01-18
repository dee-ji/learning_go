package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Number is an interface that allows both integer and float types
type Number interface {
	constraints.Integer | constraints.Float
}

// Adder takes two parameters of type Number and returns their sum
//
// Helpful information can be found here: https://www.mathisfun.com/numbers/addition.html
func Adder[T Number](a, b T) T {
	return a + b
}

func main() {
	// Test Adder with integers
	fmt.Println(Adder(2, 3))

	// Test Adder with floats
	fmt.Println(Adder(2.5, 3.5))
}
