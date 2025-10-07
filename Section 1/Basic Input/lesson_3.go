package main

import "fmt"

func main() {
	// These variables are already defined for you
	age := 25
	price := 19.99
	isAvailable := true
	productName := "Wireless Headphones"

	// TODO: Use fmt.Printf with %T to print the type of the 'age' variable
	// The output should be: "The type of age is: int"
	fmt.Printf("The type of age is: %T\n", age)
	// TODO: Use fmt.Printf with %T to print the type of the 'price' variable
	// The output should be: "The type of price is: float64"
	fmt.Printf("The type of price is: %T\n", price)
	// TODO: Use fmt.Printf with %T to print the type of the 'isAvailable' variable
	// The output should be: "The type of isAvailable is: bool"
	fmt.Printf("The type of isAvailable is: %T\n", isAvailable)
	// TODO: Use fmt.Printf with %T to print the type of the 'productName' variable
	// The output should be: "The type of productName is: string"
	fmt.Printf("The type of productName is: %T\n", productName)
}
