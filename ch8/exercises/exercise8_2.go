package main

import "fmt"

// Define a generic interface called Printable that matches a type that implements fmt.Stringer and has an underlying
// type of int or float64. Define types that meet this interface. Write a function that takes in a Printable and prints
// its value to the screen using fmt.Println.

// Define the Printable interface

type Printable[T int | float64] interface {
	fmt.Stringer
	Value() T
}

// Define a type that meets the Printable interface

type MyInt struct {
	v int
}

// Implement fmt.Stringer for MyInt

func (m MyInt) String() string {
	return fmt.Sprintf("MyInt: %d", m.v)
}

// Implement the Value method for MyInt

func (m MyInt) Value() int {
	return m.v
}

// Define another type that meets the Printable interface

type MyFloat struct {
	v float64
}

// Implement fmt.Stringer for MyFloat

func (m MyFloat) String() string {
	return fmt.Sprintf("MyFloat: %.2f", m.v)
}

// Implement the Value method for MyFloat

func (m MyFloat) Value() float64 {
	return m.v
}

// Define a generic function that works with Printable

func PrintValue[T int | float64](p Printable[T]) {
	fmt.Println(p.String())
	fmt.Printf("Underlying value: %v\n", p.Value())
}

func main() {
	// Create instances of MyInt and MyFloat
	myInt := MyInt{v: 42}
	myFloat := MyFloat{v: 3.14}

	// Use PrintValue to print the values
	PrintValue(myInt)
	PrintValue(myFloat)
}
