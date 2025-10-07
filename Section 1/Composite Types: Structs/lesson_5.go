package main

import "fmt"

// Person struct definition
type Person struct {
	Name       string
	Age        int
	IsEmployed bool
}

func main() {
	// TODO: Initialize a Person struct with name "Alice", age 28, and isEmployed true
	// Use either the field names or the struct literal syntax
	alice := Person{
		Name:       "Alice",
		Age:        28,
		IsEmployed: true,
	}

	// Print the person's information
	fmt.Printf("Name: %s\n", alice.Name)
	fmt.Printf("Age: %d\n", alice.Age)
	fmt.Printf("Employed: %t\n", alice.IsEmployed)
}
