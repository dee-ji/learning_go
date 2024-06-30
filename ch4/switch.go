package main

import "fmt"

func switch_statements() {
	// Example 4-19 The switch statement
	// words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	// for _, word := range words {
	// 	switch size := len(word); size {
	// 	case 1, 2, 3, 4:
	// 		fmt.Println(word, "is a short word!")
	// 	case 5:
	// 		wordLen := len(word)
	// 		fmt.Println(word, "is exactly the right length:", wordLen)
	// 	case 6, 7, 8, 9:
	// 	default:
	// 		fmt.Println(word, "is a long word!")
	// 	}
	// }

	// Example 4-20 The case of the missing label
loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}

	// Example 4-12 The blank switch
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}

	// Example 4-22 Rewriting a series of if statements with a blank switch
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
