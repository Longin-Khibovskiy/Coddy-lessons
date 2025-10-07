package main

import "fmt"

func main() {
	// Source slice with fruit names
	source := []string{"apple", "banana", "cherry", "date"}

	// Create a destination slice with capacity for 3 elements
	destination := make([]string, 3)

	// TODO: Use the copy function to copy elements from source to destination
	copy(destination, source)

	// Print the destination slice
	fmt.Println("Destination slice:", destination)
}
