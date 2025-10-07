package main

import "fmt"

// This function uses named return values
func getPersonInfo() (name string, age int, isStudent bool) {
	// TODO: Assign values to the named return variables
	// name should be "Alex"
	// age should be 25
	// isStudent should be true
	name = "Alex"
	age = 25
	isStudent = true
	// The "return" statement without arguments will return the named return values
	return name, age, isStudent
}

func main() {
	name, age, isStudent := getPersonInfo()
	fmt.Printf("Name: %s\nAge: %d\nStudent: %t\n", name, age, isStudent)
}
