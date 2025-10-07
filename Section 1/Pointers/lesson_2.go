package main

import "fmt"

func main() {
	// Here's a variable and a pointer to it
	originalValue := 42
	pointerToValue := &originalValue

	// TODO: Declare a new pointer variable called 'secondPointer'
	// that points to the same memory address as 'pointerToValue'
	secondPointer := &originalValue

	// Don't modify the code below
	fmt.Printf("Original value: %v\n", originalValue)
	fmt.Printf("Value through first pointer: %v\n", *pointerToValue)
	fmt.Printf("Value through second pointer: %v\n", *secondPointer)

	// Let's change the original value and see all pointers reflect the change
	originalValue = 100
	fmt.Printf("\nAfter changing original value to 100:\n")
	fmt.Printf("Original value: %v\n", originalValue)
	fmt.Printf("Value through first pointer: %v\n", *pointerToValue)
	fmt.Printf("Value through second pointer: %v\n", *secondPointer)
}
