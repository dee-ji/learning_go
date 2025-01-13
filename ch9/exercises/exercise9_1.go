package main

import (
	"errors"
	"fmt"
)

// Create a sentinel error to represent an invalid ID. In main, use errors.Is to check for the sentinel error, and print
// a message when it is found.

var ErrInvalidID = errors.New("invalid ID")

func getUserNameByID(id int) (string, error) {
	// Simulate valid and invalid IDs
	if id <= 0 {
		return "", ErrInvalidID
	}
	return fmt.Sprintf("User %d", id), nil
}

func main() {
	// Example IDs to test
	ids := []int{1, -1, 0, 5}

	for _, id := range ids {
		// Attempt to get the username by ID
		name, err := getUserNameByID(id)
		if err != nil {
			// Check if the error is ErrInvalidID
			if errors.Is(err, ErrInvalidID) {
				fmt.Printf("ID %d is invalid: %v\n", id, err)
			} else {
				fmt.Printf("Unexpected error for ID %d: %v\n", id, err)
			}
		} else {
			fmt.Printf("ID %d corresponds to username: %s\n", id, name)
		}
	}
}
