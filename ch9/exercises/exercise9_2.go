package main

import (
	"errors"
	"fmt"
)

// Define a custom error type to represent an empty field error. This error should include the name of the empty
// Employee field. In main, use errors.As to check for this error. Print out a message that includes the field name.

type EmptyFieldError struct {
	FieldName string
}

func (e *EmptyFieldError) Error() string {
	return fmt.Sprintf("the field %q is empty", e.FieldName)
}

type Employee struct {
	Name  string
	Email string
}

// ValidateEmployee checks if required fields are empty and returns a custom error if so
func ValidateEmployee(emp Employee) error {
	if emp.Name == "" {
		return &EmptyFieldError{FieldName: "Name"}
	}
	if emp.Email == "" {
		return &EmptyFieldError{FieldName: "Email"}
	}
	return nil
}

func main() {
	employees := []Employee{
		{Name: "Alice", Email: "alice@example.com"}, // Valid
		{Name: "", Email: "bob@example.com"},        // Invalid (empty Name)
		{Name: "Charlie", Email: ""},                // Invalid (empty Email)
	}

	for i, emp := range employees {
		err := ValidateEmployee(emp)
		if err != nil {
			var emptyFieldErr *EmptyFieldError
			// Check if the error is of type EmptyFieldError using errors.As
			if errors.As(err, &emptyFieldErr) {
				fmt.Printf("Employee %d: %v\n", i+1, emptyFieldErr.Error())
			} else {
				fmt.Printf("Employee %d: Unexpected error: %v\n", i+1, err)
			}
		} else {
			fmt.Printf("Employee %d: Validation passed\n", i+1)
		}
	}
}
