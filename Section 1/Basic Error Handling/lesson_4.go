package main

import (
	"fmt"
)

// This function simulates opening a file
// It returns an error if the file doesn't exist
func openFile(filename string) error {
	if filename == "data.txt" {
		return nil // nil means no error
	} else {
		return fmt.Errorf("file %s not found", filename)
	}
}

func main() {
	filename := "config.txt"

	// Call the openFile function
	err := openFile(filename)

	// TODO: Check if err is not nil (which means there was an error)
	// If there was an error, print: "Error: " followed by the error message
	// If there was no error, print: "File opened successfully"
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File opened successfully")
	}
}
