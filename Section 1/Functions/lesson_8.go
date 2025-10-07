package main

import "fmt"

// Global variable
var message string = "Hello from global scope"

func updateMessage() {
	// This creates a new local variable instead of modifying the global one
	message := "Hello from function scope"
	fmt.Println("Inside function:", message)
}

func main() {
	fmt.Println("Before function call:", message)
	updateMessage()
	message = "Hello from function scope"
	// TODO: The global message is still the original value
	// Modify the updateMessage() function to change the global variable
	fmt.Println("After function call:", message)
}
