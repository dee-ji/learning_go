package main

import "fmt"

// Write a generic function that doubles the value of any integer or float
// that's passed in to it. Define any needed generic interfaces.

func Doubler[T int | int8 | int16 | int32 | int64 | float32 | float64](d T) T {
	return d * 2
}

func main() {
	doubledNum := Doubler(5.3)
	fmt.Println(doubledNum)
}
