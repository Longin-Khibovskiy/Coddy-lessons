package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Read input
	var dimensions string
	var targetStr string
	fmt.Scanln(&dimensions)
	fmt.Scanln(&targetStr)

	// Parse dimensions
	dimParts := strings.Split(dimensions, ",")
	rows, _ := strconv.Atoi(dimParts[0])
	cols, _ := strconv.Atoi(dimParts[1])

	// Parse target number
	target, _ := strconv.Atoi(targetStr)

	// Predefined grid
	grid := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	// TODO: Write your code below
	// Use nested loops with labeled break to search for the target
	target_bul := false
search:
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == target {
				fmt.Printf("Found %d at position (%d, %d)", target, i, j)
				target_bul = true
				break search
			}
		}
	}
	if target_bul == false {
		fmt.Printf("Target %d not found", target)
	}
}
