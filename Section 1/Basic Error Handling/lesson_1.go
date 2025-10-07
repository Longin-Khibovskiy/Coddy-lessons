package main

import (
	"errors"
	"fmt"
)

// divide divides a by b and returns the result
// If b is 0, it returns an error
func divide(a, b int) (int, error) {
	// TODO: Implement the divide function
	// If b is 0, return 0 and an error with message "division by zero"
	// Otherwise, return a/b and nil for the error
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
func main() {
	// First test case
	a, b := 10, 2
	result, err := divide(a, b)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}

	// Second test case
	a, b = 8, 0
	result, err = divide(a, b)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
}
