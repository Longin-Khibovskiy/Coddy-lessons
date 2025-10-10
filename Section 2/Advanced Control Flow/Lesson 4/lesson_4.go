package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Read input
	var maxRetriesStr string
	var successAttemptStr string
	fmt.Scanln(&maxRetriesStr)
	fmt.Scanln(&successAttemptStr)

	// Parse inputs to integers
	maxRetries, _ := strconv.Atoi(maxRetriesStr)
	successAttempt, _ := strconv.Atoi(successAttemptStr)

	// Initialize attempt counter
	attempt := 1
start:
	// TODO: Write your code below
	// Use goto statement with a label to implement retry logic
	if attempt <= maxRetries {
		if attempt < successAttempt {
			fmt.Printf("Attempt %d failed\n", attempt)
			attempt++
			goto start
		}
		if attempt == successAttempt {
			fmt.Printf("Attempt %d succeeded", attempt)
		}
	}
	if attempt > maxRetries {
		fmt.Println("All attempts failed")
	}
}
