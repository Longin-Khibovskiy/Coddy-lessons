package main

import "fmt"

func main() {
	// These variables are declared but not initialized
	// Go will automatically assign them their zero values
	var number int
	var decimal float64
	var isTrue bool
	var text string

	// TODO: Print each variable to see its zero value
	// Hint: Use fmt.Println() for each variable
	fmt.Println(number)
	fmt.Println(decimal)
	fmt.Println(isTrue)
	fmt.Println(text)
}
