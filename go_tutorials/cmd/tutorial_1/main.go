package main

import (
	"errors"
	"fmt"
	// "unicode/utf8"
)

func main() {
	// var int_num int = 32767
	// fmt.Println(int_num)

	// var float_num float32 = 12345678.9
	// fmt.Println(float_num)

	// var float_num32 float32 = 10.1
	// var int_num32 int32 = 2
	// var result float32 = float_num32 + float32(int_num32)
	// fmt.Println(result)

	// var int_num1 int = 3
	// var int_num2 int = 2
	// fmt.Println(int_num1/int_num2)
	// fmt.Println(int_num1%int_num2)

	// var my_string string = "Hello" + " " + "World"
	// fmt.Println(my_string)

	// fmt.Println(utf8.RuneCountInString(my_string))

	// var my_rune rune = 'a'
	// fmt.Println(my_rune)

	// var my_bool bool = true
	// fmt.Println(my_bool)

	// my_var := "text"
	// fmt.Println(my_var)

	// var1, var2 := 1,2
	// fmt.Println(var1, var2)

	// const my_const string = "const str"
	// fmt.Println(my_const)

	// const pi float32 = 3.1415
	// fmt.Println(pi)

	print_me("Hello World")

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = int_division(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v\n", result)
	} else {
		fmt.Printf("The result of the integer division is &v with remainder %v", result)
	}
	fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
}

func print_me(print_value string) {
	fmt.Println(print_value)
}

func int_division(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
