package main

import "fmt"

func main() {
	// Product information
	productName := "Wireless Headphones"
	productPrice := 79.99
	productInStock := true
	productID := 12345

	// TODO: Use fmt.Printf() with format verbs to display the product information:
	// Product: Wireless Headphones (ID: 12345)
	// Price: $79.99
	// In Stock: true
	//
	// Hint: Use %s for strings, %d for integers, %.2f for floats, %t for booleans
	// Remember to add \n at the end of each line for proper formatting
	fmt.Printf("Product: %s (ID: %d)\nPrice: $%.2f\nIn Stock: %t", productName, productID, productPrice, productInStock)

}
