package main

import (
	"errors"
	"fmt"
	"strings"
)

// Rather than returning the first error found, return back a single error that contains all errors discovered
// during validation. Update the code in main to properly report multiple errors.

type AllEmptyFieldsError struct {
	FieldName string
}

func (e *AllEmptyFieldsError) Error() string {
	return fmt.Sprintf("the field %q is empty", e.FieldName)
}

type ValidationError struct {
	Errors []error
}

func (ve *ValidationError) Error() string {
	var messages []string
	for _, err := range ve.Errors {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, "; ")
}

type NewEmployee struct {
	Name  string
	Email string
}

// ValidateNewEmployee checks for multiple validation errors
func ValidateNewEmployee(emp NewEmployee) error {
	var validationErrors ValidationError

	if emp.Name == "" {
		validationErrors.Errors = append(validationErrors.Errors, &AllEmptyFieldsError{FieldName: "Name"})
	}
	if emp.Email == "" {
		validationErrors.Errors = append(validationErrors.Errors, &AllEmptyFieldsError{FieldName: "Email"})
	}

	if len(validationErrors.Errors) > 0 {
		return &validationErrors
	}

	return nil
}

func main() {
	// Create a list of employees to validate
	employees := []NewEmployee{
		{Name: "Alice", Email: "alice@example.com"}, // Valid
		{Name: "", Email: "bob@example.com"},        // Invalid (empty Name)
		{Name: "Charlie", Email: ""},                // Invalid (empty Email)
		{Name: "", Email: ""},                       // Invalid (both fields empty)
	}

	// Validate employees
	for i, emp := range employees {
		err := ValidateNewEmployee(emp)
		if err != nil {
			var validationErr *ValidationError
			// Check if the error is a ValidationError
			if errors.As(err, &validationErr) {
				fmt.Printf("Employee %d: Validation errors: %s\n", i+1, validationErr.Error())
			} else {
				fmt.Printf("Employee %d: Unexpected error: %v\n", i+1, err)
			}
		} else {
			fmt.Printf("Employee %d: Validation passed\n", i+1)
		}
	}
}
