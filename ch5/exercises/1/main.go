package main

import (
	"errors"
	"fmt"
)

// Code from pg 101

func add(i int, j int) (int, error) { return i + j, nil }
func sub(i int, j int) (int, error) { return i - j, nil }
func mul(i int, j int) (int, error) { return i * j, nil }
func div(i int, j int) (int, error) {
	if i == 0 || j == 0 {
		return 0, errors.New("divide by zero")
	}
	return i / j, nil
}

// The simple calculator program doesn't handle one error case: division by zero. Change the function signature for the
// math operations to return both an int and an error. In the div function, if the divisor is 0,
// return errors.New("division by zero") for the error. In all other cases, return nil.
// Adjust the main function to check for this error.
func main() {
	a := 2
	b := 0
	if result, err := add(a, b); err == nil {
		fmt.Println("Addition result:", result)
	}
	if result, err := sub(a, b); err == nil {
		fmt.Println("Subtraction result:", result)
	}
	if result, err := mul(a, b); err == nil {
		fmt.Println("Multiplication result:", result)
	}
	if result, err := div(a, b); err != nil {
		fmt.Println("Division error:", err)
	} else {
		fmt.Println("Division result:", result)
	}
}
