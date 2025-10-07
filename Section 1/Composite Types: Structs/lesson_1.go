package main

import "fmt"

// TODO: Define a custom type called 'Temperature' based on float64
type Temperature float64

func main() {
	// Create a variable of type Temperature with value 23.5
	var roomTemp Temperature = 23.5

	// Print the value and type of roomTemp
	fmt.Printf("Room temperature: %v\n", roomTemp)
	fmt.Printf("Type of roomTemp: %T\n", roomTemp)
}
