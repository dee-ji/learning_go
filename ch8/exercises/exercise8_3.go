package main

import (
	"errors"
	"fmt"
)

// Write a generic singly linked list data type. Each element can hold a comparable value and has a pointer to the next
// element in the list. The methods to implement are as follows:
// // adds a new element to the end of the linked list
// Add(T)
// // adds an element at the specified position in the linked list
// Insert(T, int)
// // returns the position of the supplied value, -1 if it's not present
// Index (T) int

// Define a generic Node

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// Define the generic SinglyLinkedList

type SinglyLinkedList[T comparable] struct {
	head *Node[T]
	size int
}

// Add adds a new element to the end of the linked list
func (list *SinglyLinkedList[T]) Add(value T) {
	newNode := &Node[T]{Value: value}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	list.size++
}

// Insert adds an element at the specified position in the linked list
func (list *SinglyLinkedList[T]) Insert(value T, position int) error {
	if position < 0 || position > list.size {
		return errors.New("position out of bounds")
	}

	newNode := &Node[T]{Value: value}
	if position == 0 {
		newNode.Next = list.head
		list.head = newNode
	} else {
		current := list.head
		for i := 0; i < position-1; i++ {
			current = current.Next
		}
		newNode.Next = current.Next
		current.Next = newNode
	}
	list.size++
	return nil
}

// Index returns the position of the supplied value, -1 if it's not present
func (list *SinglyLinkedList[T]) Index(value T) int {
	current := list.head
	position := 0
	for current != nil {
		if current.Value == value {
			return position
		}
		current = current.Next
		position++
	}
	return -1
}

// Display prints the elements of the linked list
func (list *SinglyLinkedList[T]) Display() {
	current := list.head
	for current != nil {
		fmt.Printf("%v -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	// Create a new singly linked list
	list := SinglyLinkedList[int]{}

	// Add elements to the list
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Display() // Output: 10 -> 20 -> 30 -> nil

	// Insert an element
	err := list.Insert(15, 1)
	if err != nil {
		fmt.Println(err)
	}
	list.Display() // Output: 10 -> 15 -> 20 -> 30 -> nil

	// Find the index of an element
	fmt.Println("Index of 20:", list.Index(20)) // Output: Index of 20: 2
	fmt.Println("Index of 40:", list.Index(40)) // Output: Index of 40: -1
}
