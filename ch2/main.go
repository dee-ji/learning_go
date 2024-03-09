package main

import (
	"fmt"
	"math/cmplx"
)

const x int64 = 10

const (
	id_key   = "id"
	name_key = "name"
)

const z = 20 * 10

func main() {
	complex_nums()
	type_conversions()
	int_type_converstions()

	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	// x = x + 1 // this will not compile!
	// y = "bye" // this will not compile!

	// ./main.go:27:2: cannot assign to x (neither addressable nor a map index expression)
	// ./main.go:28:2: cannot assign to y (neither addressable nor a map index expression)

	fmt.Println(x)
	fmt.Println(y)

	// Never do this!!
	__ := "double underscore"
	fmt.Println(__)

}

func complex_nums() {
	// Example 2.1: Complex numbers
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))
}

func type_conversions() {
	// Example 2.2: Type conversions
	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1, sum2)
}

func int_type_converstions() {
	// Example 2.3: Integer type conversions
	var x int = 10
	var b byte = 100
	var sum3 int = x + int(b)
	var sum4 byte = byte(x) + b
	fmt.Println(sum3, sum4)
	// Use of := to assign variables with type inference
	d := "this is a string"
	fmt.Println(d)
}
