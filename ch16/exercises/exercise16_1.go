package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// ValidateStringLength validates that string fields in the struct have lengths
// equal to or greater than the value specified in the `minStrlen` struct tag.
func ValidateStringLength(s interface{}) error {
	// Ensure the input is a struct.
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Struct {
		return errors.New("input must be a struct")
	}

	var errs []error

	// Iterate through struct fields.
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Check if the field is a string and has the `minStrlen` tag.
		if field.Kind() == reflect.String {
			minStrlenTag, ok := fieldType.Tag.Lookup("minStrlen")
			if ok {
				// Parse the tag value as an integer.
				minLength, err := strconv.Atoi(minStrlenTag)
				if err != nil {
					err := fmt.Errorf("invalid minStrlen tag value on field '%s': %v", fieldType.Name, err)
					errs = append(errs, err)
					continue
				}

				// Validate the length of the string field.
				if len(field.String()) < minLength {
					err := fmt.Errorf("field '%s' must be at least %d characters long", fieldType.Name, minLength)
					errs = append(errs, err)
				}
			}
		}
	}

	// Use errors.Join to aggregate all errors.
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

// ExampleStruct with `minStrlen` tags.
type ExampleStruct struct {
	Name    string `minStrlen:"3"`
	Address string `minStrlen:"5"`
	Age     int    // This field will be ignored.
	Comment string // No tag, so it will be ignored.
}

func main() {
	example := ExampleStruct{
		Name:    "Jo",
		Address: "NY",
		Age:     30,
		Comment: "This is a comment.",
	}

	err := ValidateStringLength(example)
	if err != nil {
		fmt.Println("Validation errors:", err)
	} else {
		fmt.Println("All fields are valid.")
	}
}
