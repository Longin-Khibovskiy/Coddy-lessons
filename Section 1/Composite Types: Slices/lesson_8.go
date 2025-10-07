package main

import "fmt"

func main() {
	// A slice of fruits
	fruits := []string{"Apple", "Banana", "Cherry", "Dragon fruit", "Elderberry"}

	// TODO: Complete the for loop using range to iterate through the fruits slice
	// Print each fruit with its position like: "1. Apple"

	for i, value := range fruits {
		// Add your code here to print each fruit with its position
		fmt.Printf("%d. %s\n", i+1, value)
	}
}
