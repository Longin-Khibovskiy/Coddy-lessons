package main

import (
	"fmt"
)

func main() {
	// Predefined username that's too short
	username := "bob"
	// Minimum required length
	minLength := 5

	// Check if username is valid
	if len(username) < minLength {
		// TODO: Create a formatted error using fmt.Errorf that includes:
		// 1. The invalid username
		// 2. The minimum required length
		// Example format: "invalid username: [username] is too short, minimum length is [minLength]"
		err := fmt.Errorf("invalid username: %s is too short, minimum length is %d", username, minLength)

		// Print the error
		fmt.Println(err)
	} else {
		fmt.Println("Username is valid!")
	}
}
