package main

import (
	"errors"
	"fmt"
)

func main() {
	// This function returns an error
	err := generateError()

	// TODO: Check if err is not nil
	// If err is not nil, print: "Error occurred: " followed by the error message
	// HINT: Use an if statement to check if err != nil
	if err != nil {
		fmt.Println("Error occurred:", err)
	}

	// This line will only execute if no error occurred
	fmt.Println("Program completed successfully")
}

// This function generates an error
func generateError() error {
	return errors.New("something went wrong")
}
