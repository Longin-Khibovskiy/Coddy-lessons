package main

import "fmt"

func main() {
	// Predefined variables to compare
	num1 := 10
	num2 := 10
	name1 := "Go"
	name2 := "Python"

	// TODO: Compare num1 and num2 using == (equality) operator
	// and store the result in numbersEqual
	numbersEqual := num1 == num2 // Replace this line

	// TODO: Compare name1 and name2 using != (inequality) operator
	// and store the result in namesNotEqual
	namesNotEqual := name1 != name2 // Replace this line

	// Print the results
	fmt.Println("Are the numbers equal?", numbersEqual)
	fmt.Println("Are the names different?", namesNotEqual)
}
