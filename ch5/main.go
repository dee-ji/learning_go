package main

import (
	"fmt"
	// "log"
	"os"
)

func main() {
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

	res, rem, er := nakedDivAndRemainder(5, 2)
	if er != nil {
		fmt.Println(er)
		os.Exit(1)
	}
	fmt.Println(res, rem)

	myClosure()
	myPerson()
	cat()
	deferExample()

	// f, closer, err := getFile(os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer closer()
}
