package main

import (
	"fmt"
	"os"
)

// Write a function called fileLen that has an input parameter of type string and returns an int and an error.
// The function takes in a filename and returns the number of bytes in the file. If there is an error reading the file,
// return the error. Use defer to make sure the file is closed properly.
func fileLen(filename string) (int, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// Ensure the file is closed properly
	defer file.Close()
	// Get the file info to retrieve the size
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	// Return the size of the file in bytes
	return int(fileInfo.Size()), nil
}

func main() {
	filename := "example.txt"

	length, err := fileLen(filename)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("The file '%s' has %d bytes.\n", filename, length)
	}
}
