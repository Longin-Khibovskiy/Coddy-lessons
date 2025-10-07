package main

import "fmt"

func main() {
	// A slice of fruits
	fruits := []string{"apple", "banana", "orange"}

	// TODO: Append "grape" and "kiwi" to the fruits slice
	// Write your code here
	fruits = append(fruits, "grape", "kiwi")
	// Print the updated slice
	fmt.Println("My fruit collection:", fruits)
}
