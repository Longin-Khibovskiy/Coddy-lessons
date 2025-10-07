package main

import "fmt"

func main() {
	// Given variables
	name := "Gopher"
	age := 5
	height := 0.75
	isActive := true

	// TODO: Use fmt.Printf to print "Name: Gopher" using the name variable and the %s format specifier
	fmt.Printf("Name: %s\n", name)
	// TODO: Use fmt.Printf to print "Age: 5 years" using the age variable and the %d format specifier
	fmt.Printf("Age: %d years\n", age)
	// TODO: Use fmt.Printf to print "Height: 0.75 meters" using the height variable and the %.2f format specifier
	fmt.Printf("Height: %.2f meters\n", height)
	// TODO: Use fmt.Printf to print "Active: true" using the isActive variable and the %t format specifier
	fmt.Printf("Active: %t\n", isActive)
	// TODO: Use fmt.Printf to print "Name type: string" using the name variable and the %T format specifier
	fmt.Printf("Name type: %T\n", name)
}
