package main

import "fmt"

// TODO: Complete the greet function that takes a name parameter
// and returns a greeting string in the format: "Hello, " + name + "!"
// Example: if name is "Gopher", return "Hello, Gopher!"
func greet(name string) string {
	// Your code here
	return "Hello, " + name + "!"
}

func main() {
	// Test the function with a predefined name
	name := "Gopher"
	message := greet(name)
	fmt.Println(message)

	// Expected output: Hello, Gopher!
}
