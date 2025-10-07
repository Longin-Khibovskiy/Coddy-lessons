package main

import "fmt"

// Person struct with name and age fields
type Person struct {
	name string
	age  int
}

func main() {
	// Create a Person struct
	person := Person{
		name: "John",
		age:  25,
	}

	// Create a pointer to the Person struct
	personPtr := &person

	// Print the original person information
	fmt.Println("Original person:")
	fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)

	// TODO: Use the pointer to change the person's name to "Jane" and age to 30
	// Hint: You can access struct fields through a pointer using the dot notation
	personPtr.name = "Jane"
	personPtr.age = 30

	// Print the updated person information
	fmt.Println("Updated person:")
	fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)
}
