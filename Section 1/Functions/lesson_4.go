package main

import "fmt"

// TODO: Add parameters to this function to accept a greeting and a name
// The function should take two string parameters: greeting and name
// You can write them as (greeting string, name string) or (greeting, name string)
func createGreeting(greeting string, name string) string {
	// TODO: Return a string that combines the greeting and name
	// For example: "Hello, John!"
	// Use the format: greeting + ", " + name + "!"
	return greeting + ", " + name + "!"
}

func main() {
	// These test calls are already defined for you
	message := createGreeting("Hello", "Gopher")
	fmt.Println(message)

	// Test with different values
	message = createGreeting("Welcome", "Friend")
	fmt.Println(message)
}
