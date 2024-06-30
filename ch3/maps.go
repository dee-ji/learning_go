package main

import "fmt"
import "maps"

func maps_examples() {
	fmt.Println("#######################")
	fmt.Println("Starting Maps section")
	fmt.Println("#######################")
	// The map type is written as map[keyType]vauleType
	// This example would be a map with strings as keys and integers as values

	// var nilMap map[string]int

	// However, this nilMap value is nil and len of 0. Writting to this map will cause a panic!
	// Use := notation to define maps

	// totalWins := map[string]int{}

	// None empty map definition
	teams := map[string][]string {
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	// The VSCode interpetor for Go is saying that []string before the values is redundant. So it's an estetic to use or not here.

	fmt.Println(teams)
	fmt.Println("This map has size", len(teams))
	// Using the make function you can create map with a default size. This map can still grow past this size.
	ages := make(map[int][]string, 10)
	fmt.Println(ages)
	fmt.Println("This map has size", len(ages))

	// Reading and writing to a map
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])

	// The comma ok idiom
	// m := map[string]int{
	// 	"hello": 5,
	// 	"world": 0,
	// }
	// v, ok := m["hello"]
	// fmt.Println(v, ok)

	// v, ok = m["world"]
	// fmt.Println(v, ok)

	// v, ok = m["goodbye"]
	// fmt.Println(v, ok)

	// Deleting maps
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	// delete(m, "hello")
	// fmt.Println(m)

	// Emptying maps
	// fmt.Println(m, len(m))
	// clear(m)
	// fmt.Println(m, len(m))

	// Comparing Maps
	n := map[string]int{
		"hello": 5,
		"world": 10,
	}
	fmt.Println("Are maps m and n the same?")
	fmt.Println(maps.Equal(m, n))

	// Usings maps as sets
	// Go does not have a set data type but you can use maps to simulate their behavior
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println("Simlulating a 'set' in Go with a map")
	fmt.Println(intSet)
	fmt.Println(vals)
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	// Struct implementation of a set in Go
	fmt.Println("Simlulating a 'set' in Go with a struct value")
	intStructSet := map[int]struct{}{}
	values := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range values {
		intStructSet[v] = struct{}{}
	}
	fmt.Println(intStructSet)
	fmt.Println(values)
	fmt.Println(len(values), len(intStructSet))
	fmt.Println(intStructSet[5])
	fmt.Println(intStructSet[500])
	if _, ok := intStructSet[5]; ok {
		fmt.Println("5 is in the set")
	}

	// Structs
	fmt.Println("Here are examples of Structs")
	type person struct {
		name string
		age int
		pet string
	}

	var fred person

	bob := person{}

	julia := person{
		"Julia",
		40,
		"cat",
	}

	beth := person{
		age: 30,
		name: "Beth",
	}
	fmt.Println(fred)
	fmt.Println(bob)
	fmt.Println(julia)
	fmt.Println(beth)

	// Anonymous structs
	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(pet)

}