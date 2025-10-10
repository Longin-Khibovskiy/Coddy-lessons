package main

import "fmt"

func main() {
	// Read input
	var alertLevel string
	fmt.Scanln(&alertLevel)

	// TODO: Write your code below
	// Create a switch statement with fallthrough for the alert system
	switch alertLevel {
	case "critical":
		fmt.Println("CRITICAL: System shutdown imminent!")
		fallthrough
	case "high":
		fmt.Println("HIGH: Immediate attention required!")
		fallthrough
	case "medium":
		fmt.Println("MEDIUM: Please review when possible")
	case "low":
		fmt.Println("LOW: Informational only")
	}
}
