package main

import "fmt"

// Person struct definition
type Person struct {
	name      string
	age       int
	isStudent bool
}

func main() {
	// TODO: Create a new Person struct instance with
	// name: "Alice", age: 25, isStudent: true
	person := Person{"Alice", 25, true}
	// Don't modify the code below
	fmt.Printf("Name: %s\n", person.name)
	fmt.Printf("Age: %d\n", person.age)
	fmt.Printf("Is Student: %t\n", person.isStudent)
}
