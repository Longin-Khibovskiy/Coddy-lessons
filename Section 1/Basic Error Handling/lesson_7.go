package main

import (
	"errors"
	"fmt"
)

// This function divides x by y and returns an error if y is zero
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}

func main() {
	// Test values
	numerator := 10.0
	denominator := 0.0

	// Call the divide function
	result, err := divide(numerator, denominator)

	// TODO: Check if err is not nil (there is an error)
	// If there is an error, print it using fmt.Println(err)
	// Otherwise, print the result using fmt.Printf("Result: %.2f\n", result)
	if err != nil {
		fmt.Println(err)
	} else {
		// Remove this line and implement proper error handling above
		fmt.Printf("Result: %.2f\n", result)
	}
}
