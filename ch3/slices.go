package main

import "fmt"

func main() {
	// Example 3.1 Understanding capacity
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

	// Example 3.2 Declaring a slice that might stay nil
	// var data []int
	// Example 3.3 Declaring a slice with default values
	// data := []int{2, 4, 6, 8}  // numbers we appreciate

	// Example 3.4 Slicing slices
	w := []string{"a", "b", "c", "d"}
	y := w[:2]
	z := w[1:]
	d := w[1:3]
	e := w[:]
	fmt.Println("w:", w)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	// Example 3.5 Slices with overlapping storage
	w[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("w:", w)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// Example 3.6 Append makes overlapping slices more confusing
	h := []string{"a", "b", "c", "d"}
	i := h[:2]
	fmt.Println(i)
	fmt.Println(cap(h), cap(i))
	i = append(i, "z")
	fmt.Println("h:", h)
	fmt.Println("i:", i)

	// Example 3.7 Even more confusing slices
	j := make([]string, 0, 5)
	j = append(j, "a", "b", "c", "d")
	k := j[:2]
	l := j[2:]
	fmt.Println(cap(j), cap(k), cap(l))
	k = append(k, "i", "j", "k")
	j = append(j, "x")
	l = append(l, "y")
	fmt.Println("j:", j)
	fmt.Println("k:", k)
	fmt.Println("l:", l)

	// Example 3.8 The full slice expression protects against append
	m := make([]string, 0, 5)
	m = append(m, "a", "b", "c", "d")
	n := m[:2:2]
	o := m[2:4:4]
	n = append(n, "i", "j", "k")
	m = append(m, "x")
	o = append(o, "y")
	fmt.Println("m:", m)
	fmt.Println("n:", n)
	fmt.Println("o:", o)

	// Using copy for a slice
	p := []int{1, 2, 3, 4}
	q := make([]int, 4)
	num := copy(q, p)
	fmt.Println(q, num)
	// Copy subset of the slice
	r := make([]int, 2)
	num1 := copy(r, p)
	fmt.Println(r, num1)
	// Copy from the middle of the slice
	s := make([]int, 2)
	num2 := copy(s, p[2:])
	fmt.Println(s, num2)
	// Copy between 2 slices that cover overlapping sections
	num3 := copy(p[:3], p[1:])
	fmt.Println(p, num3)
}