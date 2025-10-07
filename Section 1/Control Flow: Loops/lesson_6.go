package main

import "fmt"

func main() {
	// Number of rows and columns in our grid
	rows := 3
	columns := 3

	// TODO: Create a nested loop to print a grid of asterisks (*)
	// The outer loop should iterate through the rows
	// The inner loop should iterate through the columns
	for i := 0; i < rows; i++ {
		// Add your inner loop here to print asterisks in each column
		for j := 1; j <= columns; j++ {
			fmt.Printf("* ")
		}
		// This prints a new line after each row
		fmt.Println()
	}
}
