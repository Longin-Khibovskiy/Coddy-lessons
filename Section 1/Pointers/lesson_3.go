package main

import "fmt"

func main() {
	// This is our sample variable
	name := "Gopher"

	// This declares a pointer variable to store the address of a string
	var namePointer *string

	// TODO: Use the address-of operator (&) to get the memory address of 'name'
	// and store it in the 'namePointer' variable
	namePointer = &name
	// Print the value that the pointer points to
	fmt.Printf("The value at that memory address is: %v\n", *namePointer)
}
