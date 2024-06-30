package main

import "fmt"

func string_examples() {
	fmt.Println("#######################")
	fmt.Println("Starting String section")
	fmt.Println("#######################")
	// Slicing a string
	// var s string = "Hello there"
	// var s2 string = s[4:7]
	// var s3 string = s[:5]
	// var s4 string = s[6:]

	// fmt.Println(len(s))

	// Type conversion bewteen rune, string, and byte
	// var a rune = 'x'
	// var s string = string(a)
	// var b byte = 'y'
	// var s2 string = string(b)

	// Converting int to string should not be done!
	// var x int = 65
	// var y string = string(x)
	// This will yield "A" not "65" and when running the make build go fmt it will not allow this to compile
	// fmt.Println(y)

	// Example 3-9 Conveting strings to slices
	var s string = "Hello, ⚪"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)

	// fmt.Println(s)
	// fmt.Println(s2)
	// fmt.Println(s3)
	// fmt.Println(s4)

	fmt.Println("#######################")
}