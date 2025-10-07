package main

import "fmt"

func main() {
	// These variables are already set up for you
	a := true
	b := false
	c := true

	// TODO: Add parentheses to the expression below to make it evaluate to true
	// HINT: The expression needs to evaluate to true
	result := a && !b || c

	// Don't modify the line below
	fmt.Println(result)
}
