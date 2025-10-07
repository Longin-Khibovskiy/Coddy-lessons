package main

import "fmt"

func main() {
	// TODO: Create a slice of strings with length 3 and capacity 5 using make
	// var names = ...
	names := make([]string, 3, 5)
	// Assign values to the slice
	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Charlie"

	// Print the slice
	fmt.Println("Names:", names)

	// Print the length and capacity of the slice
	fmt.Printf("Length: %d, Capacity: %d\n", len(names), cap(names))
}
