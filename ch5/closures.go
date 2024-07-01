package main

import "fmt"

func myClosure() {
	a := 20
	f := func() {
		fmt.Println(a)
		// using := versus = creates a new a that ceases to exist when the closure exits
		a := 30
		// a = 30
		fmt.Println(a)
	}
	f()
	fmt.Println(a)
}
