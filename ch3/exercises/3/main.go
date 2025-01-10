package main

import "fmt"

// Write a program that defines a struct called Employee with three fields:
// - firstName
// - lastName
// - id
// The first two fields are of type string, and the last field (id) is of type int. Create three instances of this
// struct using whatever values you'd like. Initialize the first one using the struct literal style without names,
// the second using the struct literal style with names, and the third with a var declaration. Use dot notation to
// populate the fields in the third struct. Print out all three structs.
func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	// struct literal definition
	david := Employee{"David", "Jones", 1}
	// struct literal with variables defined
	kate := Employee{
		firstName: "Kate",
		lastName:  "Jones",
		id:        2,
	}
	// var declaration
	var sam Employee
	sam.firstName = "Sam"
	sam.lastName = "Jones"
	sam.id = 3

	// Print out all the values
	fmt.Println("Employee 1:", david)
	fmt.Println("Employee 2:", kate)
	fmt.Println("Employee 3:", sam)
}
