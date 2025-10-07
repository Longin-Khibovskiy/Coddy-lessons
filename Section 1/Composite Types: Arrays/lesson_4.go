package main

import "fmt"

func main() {
	// This is our array of fruits
	fruits := [5]string{"apple", "banana", "cherry", "date", "elderberry"}

	// TODO: Create a variable called firstFruit and assign it the first element of the fruits array
	firstFruit := fruits[0]
	// TODO: Create a variable called thirdFruit and assign it the third element of the fruits array
	thirdFruit := fruits[2]
	// Print the fruits
	fmt.Println("The first fruit is:", firstFruit)
	fmt.Println("The third fruit is:", thirdFruit)
}
