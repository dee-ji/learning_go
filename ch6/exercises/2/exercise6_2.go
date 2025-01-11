package main

import "fmt"

// Write two functions. The UpdateSlice function takes in a []string and a string. It sets the last position in the
// passed-in slice to the passed-in string. At the end of UpdatedSlice, print the slice after making the change. The
// GrowSlice function also takes in a []string and a string. It appends the string onto the slice. At the end of
// GrowSlice, print the slice after making the change. Called these functions from main. Print out the slice before each
// function is called and after each function is called. Do you understand why some changes are visible in main and why
// some changes are not?

func UpdateSlice(inputSlice []string, word string) []string {
	if len(inputSlice) > 0 {
		inputSlice[len(inputSlice)-1] = word
	}
	fmt.Println("Inside UpdateSlice: ", inputSlice)
	return inputSlice
}

func GrowSlice(inputSlice []string, word string) []string {
	inputSlice = append(inputSlice, word)
	fmt.Println("Inside GrowSlice: ", inputSlice)
	return inputSlice
}
func main() {
	// Initial slice
	mySlice := []string{"a", "b", "c"}

	// Print the slice before UpdateSlice
	fmt.Println("Before UpdateSlice:", mySlice)

	// Call UpdateSlice
	UpdateSlice(mySlice, "z")

	// Print the slice after UpdateSlice
	fmt.Println("After UpdateSlice:", mySlice)

	// Print the slice before GrowSlice
	fmt.Println("Before GrowSlice:", mySlice)

	// Call GrowSlice
	GrowSlice(mySlice, "d")

	// Print the slice after GrowSlice
	fmt.Println("After GrowSlice:", mySlice)

	// Reasoning: Since GrowSlice operates on a copy of the slice header, the changes are not reflected in main.
	// Correct way is to return the updated copy of the slice
	// myUpdatedSlice := GrowSlice(mySlice, "d")
	// fmt.Println("After GrowSlice:", myUpdatedSlice)

	// Output:
	// Before UpdateSlice: [a b c]
	// Inside UpdateSlice:  [a b z]
	// After UpdateSlice: [a b z]
	// Before GrowSlice: [a b z]
	// Inside GrowSlice:  [a b z d]
	// After GrowSlice: [a b z]
}
