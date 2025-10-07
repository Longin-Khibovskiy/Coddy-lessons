package main

import "fmt"

func main() {
	// Some predefined variables
	age1 := 25
	age2 := 30
	age3 := 25

	// TODO: Complete these comparison expressions using <, >, <=, or >=
	isAge1LessThanAge2 := age1 < age2              // Check if age1 is less than age2
	isAge2GreaterThanAge1 := age2 > age1           // Check if age2 is greater than age1
	isAge1LessThanOrEqualToAge3 := age1 <= age3    // Check if age1 is less than or equal to age3
	isAge2GreaterThanOrEqualToAge3 := age2 >= age3 // Check if age2 is greater than or equal to age3

	// Print the results
	fmt.Println("Is age1 less than age2?", isAge1LessThanAge2)
	fmt.Println("Is age2 greater than age1?", isAge2GreaterThanAge1)
	fmt.Println("Is age1 less than or equal to age3?", isAge1LessThanOrEqualToAge3)
	fmt.Println("Is age2 greater than or equal to age3?", isAge2GreaterThanOrEqualToAge3)
}
