package main

import "fmt"

func main() {
	// TODO: Initialize an array called 'favoriteNumbers' with these 5 integers: 7, 42, 8, 13, 99
	// You can use either syntax:
	// favoriteNumbers := [5]int{7, 42, 8, 13, 99}
	// OR
	// var favoriteNumbers [5]int = [5]int{7, 42, 8, 13, 99}
	favoriteNumbers := [...]int{7, 42, 8, 13, 99}
	// This will print your array
	fmt.Printf("My favorite numbers are: %v\n", favoriteNumbers)
}
