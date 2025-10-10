package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Read input
	var rowsInput string
	var skipConditionInput string
	fmt.Scanln(&rowsInput)
	fmt.Scanln(&skipConditionInput)

	// Parse inputs
	numRows, _ := strconv.Atoi(rowsInput)
	skipCondition, _ := strconv.Atoi(skipConditionInput)

	// Provided 2D data array
	data := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	// TODO: Write your code below
	// Use nested loops with labeled continue to process the data
	// Check each row for the skip condition number
	// Print appropriate messages based on whether the condition is found
outer:
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == skipCondition {
				fmt.Printf("Skipping row %d due to condition", i)
				continue outer
			}
		}
		fmt.Printf("Processing row %d:", i)
		for _, value := range data[i] {
			fmt.Printf(" %d", value)
		}
		fmt.Println()
	}
}
