package main

import "fmt"

func main() {
	// A slice of fruits
	fruits := []string{"apple", "banana", "cherry", "date", "elderberry"}

	// TODO: Access the first fruit (index 0) and store it in a variable called firstFruit
	firstFruit := fruits[0]
	// TODO: Access the third fruit (index 2) and store it in a variable called thirdFruit
	thirdFruit := fruits[2]
	// Print the results
	fmt.Println("The first fruit is:", firstFruit)
	fmt.Println("The third fruit is:", thirdFruit)
}
