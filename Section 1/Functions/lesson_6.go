package main

import "fmt"

// This function should return three values: name, age, and isStudent
func getPersonInfo() (string, int, bool) {
	// Variables are already defined for you
	name := "Alex"
	age := 25
	isStudent := true

	// TODO: Return all three values (name, age, isStudent)
	return name, age, isStudent
}

func main() {
	// Call the function and store the returned values
	name, age, isStudent := getPersonInfo()

	// Print the values
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Is Student: %t\n", isStudent)
}
