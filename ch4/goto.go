package main

import (
	"fmt"
	"math/rand"
)

func goto_statements() {
	// This doesn't work!!!
	// .\goto.go:7:7: goto skip jumps over declaration of b at .\goto.go:8:4
	//.\goto.go:13:8: goto inner jumps into block starting at .\goto.go:15:11
	// 	a := 0
	// 	goto skip
	// 	b := 20
	// skip:
	// 	c := 30
	// 	fmt.Println(a, b, c)
	// 	if c > a {
	// 		goto inner
	// 	}
	// 	if a < b {
	// 	inner:
	// 	fmt.Println("a is less than b")
	// 	}
	// Example 4-24 A reason to use goto
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)
}
