package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

// pointer receiver

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// value receiver

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong: ", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight: ", c.String())
}

func main() {
	// Regular example with variable declaration
	var c Counter
	fmt.Println(c.String())
	// Go will automatically takes the address of the local variable
	// This could be viewed as being written as: (&c).Increment()
	c.Increment()
	fmt.Println(c.String())

	// Example with pointer variable
	counter := &Counter{}
	// If you call a value receiver on a pointer variable,
	// Go automatically dereferences the pointer when calling the method
	// This could be viewed as being written as: (*c).String()
	fmt.Println(counter.String())
	counter.Increment()
	fmt.Println(counter.String())

	var c2 Counter
	doUpdateWrong(c2)
	fmt.Println("in main: ", c2.String())
	doUpdateRight(&c2)
	fmt.Println("in main: ", c2.String())
}
