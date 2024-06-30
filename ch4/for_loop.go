package main

import "fmt"

func for_loop() {
	// Example 4.8 A complete for statement
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Leaving off the initialization
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// Leaving off the increment
	for i := 0; i < 10; {
		fmt.Println(i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}

	// Example 4-9 A condition-only for statement
	// i := 1
	// for i < 100 {
	// 	fmt.Println(i)
	// 	i = i * 2
	// }

	// Example 4-13 The for-range loop
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	// Example 4-14 Ignoring the slice index for a for-range loop
	// evenVals := []int{2,4,6,8,10,12}
	// for _, v := range evenVals {
	// 	fmt.Println(v)
	// }

	// 4-15 Map iteration order varies
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop,", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	// Example 4-16 Iterating over strings
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	// Example 4-17 Modifying the value doesn't modify the source
	// evenVals := []int{2,4,6,8,10,12}
	// for _, v := range evenVals {
	// 	v *= 2
	// }
	// fmt.Println(evenVals)

	// Example 4-18 Labels in for loops
	// samples := []string{"hello", "apple_π!"}
	// outer:
	// 	for _, sample := range samples {
	// 		for i, r := range sample {
	// 			fmt.Println(i, r, string(r))
	// 			if r == 'l' {
	// 				continue outer
	// 			}
	// 		}
	// 		fmt.Println()
	// }
}
