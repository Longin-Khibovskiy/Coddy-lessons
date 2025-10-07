package main

import "fmt"

func main() {
	// Here's our original slice of fruits
	fruits := []string{"apple", "banana", "orange", "grape", "kiwi"}

	// TODO: Create a new slice called 'firstThree' containing only the first three fruits
	firstThree := fruits[:3]
	// TODO: Create a new slice called 'lastTwo' containing only the last two fruits
	lastTwo := fruits[3:]
	// Print the results
	fmt.Println("First three fruits:", firstThree)
	fmt.Println("Last two fruits:", lastTwo)
}
