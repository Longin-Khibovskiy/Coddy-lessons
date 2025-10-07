package main

import "fmt"

// Address struct contains location information
type Address struct {
	Street  string
	City    string
	ZipCode string
}

// Person struct should embed the Address struct
type Person struct {
	Name string
	Age  int
	Address
	// TODO: Embed the Address struct here (just one line)
}

func main() {
	// Create a new person with address information
	person := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "Wonderland",
			ZipCode: "12345",
		},
	}

	// Print person information including address
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)

	// TODO: Print the address fields directly from the person instance
	// (without using person.Address.Street, etc.)
	fmt.Println("Street:", person.Street)
	fmt.Println("City:", person.City)
	fmt.Println("ZipCode:", person.ZipCode)
}
