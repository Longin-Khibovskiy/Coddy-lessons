package main

import "fmt"

func main() {
	// Outer scope variable
	message := "Hello from outer scope"

	isLoggedIn := true

	if isLoggedIn {
		// TODO: Create a new variable named 'message' with the value "Hello from inner scope"
		message := "Hello from inner scope"
		// This prints the inner scope message
		fmt.Println(message)
	}

	// This prints the outer scope message
	fmt.Println(message)
}
