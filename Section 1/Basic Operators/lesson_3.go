package main

import "fmt"

func main() {
	// These variables are already defined for you
	number := 17
	divisor := 5

	// TODO: Calculate the remainder when 'number' is divided by 'divisor'
	// using the modulo operator (%) and store it in the 'remainder' variable
	remainder := number % divisor

	// This will print the result
	fmt.Println("The remainder when", number, "is divided by", divisor, "is:", remainder)
}
