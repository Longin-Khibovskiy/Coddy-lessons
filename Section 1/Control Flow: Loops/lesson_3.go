package main

import "fmt"

func main() {
	// These variables are already set up for you
	keepRunning := true
	count := 0

	// TODO: Complete the for loop that uses only a condition (like a while loop)
	// The loop should continue as long as keepRunning is true
	for count <= 4 {
		if keepRunning == true {
			// Print the current count
			fmt.Println(count)

			// Increment count
			count++

			// Check if count is 5, and if so, set keepRunning to false
			if count >= 5 {
				keepRunning = false
			}
		}
	}
}
