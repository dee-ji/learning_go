package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Embed the rights files from the ./uhdr_docs directory
//
//go:embed uhdr_docs/*
var rightsFiles embed.FS

func main() {
	// Base path for embedded files
	basePath := "uhdr_docs"

	// Map language to file name
	languageMap := map[string]string{
		"english":  "english_rights.txt",
		"japanese": "japanese_rights.txt",
		"russian":  "russian_rights.txt",
	}

	// Check if the correct number of arguments is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: exercise11 <language>")
		fmt.Println("Available languages:")
		for lang := range languageMap {
			fmt.Printf("  - %s\n", lang)
		}
		return
	}

	// Get the language from the command-line arguments
	language := strings.ToLower(os.Args[1])

	// Find the corresponding file for the language
	fileName, exists := languageMap[language]
	if !exists {
		fmt.Printf("Error: Unsupported language '%s'.\n", language)
		fmt.Println("Available languages:")
		for lang := range languageMap {
			fmt.Printf("  - %s\n", lang)
		}
		return
	}

	// Construct the full path for the embedded file
	embeddedPath := filepath.Join(basePath, fileName)

	// Read and print the embedded file contents
	content, err := rightsFiles.ReadFile(embeddedPath)
	if err != nil {
		fmt.Printf("Error reading embedded file for language '%s': %v\n", language, err)
		return
	}

	fmt.Printf("Contents of %s:\n", fileName)
	fmt.Println(string(content))
}
