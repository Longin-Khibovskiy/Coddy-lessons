package main

import "fmt"

func main() {
	// Here's our string variable
	message := "Hello, pointers!"

	// This is a pointer to our message variable
	messagePtr := &message

	// TODO: Dereference the messagePtr to get the value it points to
	// and store it in the variable 'value'
	value := *messagePtr

	// Print the results
	fmt.Println("The pointer points to the value:", value)
}
