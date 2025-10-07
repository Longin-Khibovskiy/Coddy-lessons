package main

import (
	"errors"
	"fmt"
)

// celsiusToFahrenheit converts Celsius to Fahrenheit
// Returns an error if the temperature is below absolute zero (-273.15°C)
func celsiusToFahrenheit(celsius float64) (float64, error) {
	// TODO: Implement the conversion logic
	// 1. Check if temperature is below absolute zero (-273.15°C)
	// 2. If valid, convert using the formula: F = C × 9/5 + 32
	// 3. Return appropriate value and error
	if celsius < -273.15 {
		return 0, errors.New("Error: temperature below absolute zero")
	}
	return celsius*(9.0/5.0) + 32, nil
}

func main() {
	// Test valid temperature
	fmt.Println("Converting 25°C to Fahrenheit:")
	fahrenheit, err := celsiusToFahrenheit(25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("25°C = %.2f°F\n", fahrenheit)
	}

	// Test another valid temperature
	fmt.Println("\nConverting 0°C to Fahrenheit:")
	fahrenheit, err = celsiusToFahrenheit(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("0°C = %.2f°F\n", fahrenheit)
	}

	// Test invalid temperature (below absolute zero)
	fmt.Println("\nConverting -300°C to Fahrenheit:")
	fahrenheit, err = celsiusToFahrenheit(-300)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("-300°C = %.2f°F\n", fahrenheit)
	}
}
