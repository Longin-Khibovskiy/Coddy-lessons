package main

import "fmt"

func main() {
	// TODO: Write a for loop that prints numbers from 1 to 5
	// Each number should be on its own line
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	// TODO: Print a blank line for separation
	fmt.Println()
	// TODO: Write a for loop that calculates the sum of numbers from 1 to 5
	// Store the result in a variable named 'sum'
	// After the loop, print "Sum: 15"
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Printf("Sum: %d\n", sum)
	// TODO: Print a blank line for separation
	fmt.Println()
	// TODO: Write a for loop that starts at 10 and counts down to 1
	// Each number should be on its own line
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}
}
