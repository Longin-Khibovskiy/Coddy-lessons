package main

import "fmt"

func main() {
	// This is our array of favorite fruits
	favoriteFruits := [5]string{"Apple", "Banana", "Orange", "Mango", "Strawberry"}

	// TODO: Use the len() function to find the length of the favoriteFruits array
	// and store it in the numberOfFruits variable
	var numberOfFruits int
	numberOfFruits = len(favoriteFruits)
	// This will print the number of fruits in the array
	fmt.Printf("There are %d fruits in the array\n", numberOfFruits)
}
