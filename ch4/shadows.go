package main

import (
	"fmt"
	"math/rand"
)

func shadows() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)

	// Example 4-5 if and else
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	// Example 4-6 Scoping a variable at an if statement
	if m := rand.Intn(10); m == 0 {
		fmt.Println("That's too low")
	} else if m > 5 {
		fmt.Println("That's too big:", m)
	} else {
		fmt.Println("That's a good number:", m)
	}

	// Example 4-7 Out of scope...
	if o := rand.Intn(10); o == 0 {
		fmt.Println("That's too low")
	} else if o > 5 {
		fmt.Println("That's too big:", o)
	} else {
		fmt.Println("That's a good number:", o)
	}
	// This would be undefined
	// # go_blocks_shadows_control_structures
	// .\shadows.go:44:14: undefined: o
	// fmt.Println(o)
}
