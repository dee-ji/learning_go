package main

import "fmt"

// Create a struct named Person with three fields: FirstName and LastName of type string and Age of type int. Write a
// function called MakePerson that takes in firstName, lastName, and age and returns a Person. Write a second function
// MakePersonPointer that takes in firstName, lastName, and age and returns *Person. Call both from main. Compile your
// program with go build -gcflags="-m". This both compiles your code and prints out which values escape the heap. Are
// you surprised about what escapes?

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	john := MakePerson("John", "Doe", 30)
	jane := MakePersonPointer("Jane", "Smith", 25)

	fmt.Println(john)
	fmt.Println(jane)
}

// Output
// $ go build -gcflags="-m"
// ./exercise6_1.go:17:6: can inline MakePerson
// ./exercise6_1.go:25:6: can inline MakePersonPointer
// ./exercise6_1.go:34:20: inlining call to MakePerson
// ./exercise6_1.go:35:27: inlining call to MakePersonPointer
// ./exercise6_1.go:37:13: inlining call to fmt.Println
// ./exercise6_1.go:38:13: inlining call to fmt.Println
// ./exercise6_1.go:17:17: leaking param: firstName to result ~r0 level=0
// ./exercise6_1.go:17:35: leaking param: lastName to result ~r0 level=0
// ./exercise6_1.go:25:24: leaking param: firstName
// ./exercise6_1.go:25:42: leaking param: lastName
// ./exercise6_1.go:26:9: &Person{...} escapes to heap
// ./exercise6_1.go:35:27: &Person{...} escapes to heap
// ./exercise6_1.go:37:13: ... argument does not escape
// ./exercise6_1.go:37:14: john escapes to heap
// ./exercise6_1.go:38:13: ... argument does not escape
