package main

import "fmt"

func main() {
	// Array of fruits
	fruits := [5]string{"Apple", "Banana", "Orange", "Grape", "Mango"}

	// TODO: Complete the for loop to print each fruit in the array
	for i := 0; i < len(fruits); i++ {
		// Add your code here to print each fruit
		fmt.Println(fruits[i])
	}
}
