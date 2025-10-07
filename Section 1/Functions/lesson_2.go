package main

import "fmt"

// TODO: Complete the greet function that takes a name parameter
// and returns a greeting string

func greet(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	// These test cases are already set up for you
	name1 := "Alice"
	name2 := "Bob"

	// Testing the greet function
	message1 := greet(name1)
	message2 := greet(name2)

	fmt.Println(message1)
	fmt.Println(message2)
}
