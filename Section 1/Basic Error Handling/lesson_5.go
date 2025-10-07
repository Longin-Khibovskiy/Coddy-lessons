package main

import (
	"errors"
	"fmt"
)

func main() {
	// Test with a valid username
	result, err := validateUsername("gopher123")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// Test with an invalid username
	result, err = validateUsername("go")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

func validateUsername(username string) (string, error) {
	if len(username) < 5 {
		// TODO: Return an error with the message "username must be at least 5 characters long"
		// Hint: Use errors.New() to create a new error
		return "", errors.New("username must be at least 5 characters long") // Replace this line
	}
	return "Username is valid", nil
}
