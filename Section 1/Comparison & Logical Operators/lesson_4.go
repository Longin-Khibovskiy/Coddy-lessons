package main

import "fmt"

func main() {
	// These variables track whether assignments are complete
	homeworkComplete := false
	extraCreditComplete := true

	// TODO: Use the logical OR operator (||) to determine if the student passes
	// A student passes if either homework OR extra credit is complete
	passesClass := homeworkComplete || extraCreditComplete

	// Display the result
	fmt.Println("Homework complete:", homeworkComplete)
	fmt.Println("Extra credit complete:", extraCreditComplete)
	fmt.Println("Student passes class:", passesClass)
}
