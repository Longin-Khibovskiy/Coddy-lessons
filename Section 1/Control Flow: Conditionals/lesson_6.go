package main

import "fmt"

func main() {
	// This variable represents a day of the week
	day := "Saturday"

	// TODO: Complete the switch statement to check if the day is a weekday or weekend
	switch day {
	case "Sunday":
		fmt.Println("It's the weekend!")
	case "Saturday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
		// Add cases for weekend days (Saturday and Sunday) and print "It's the weekend!"

		// Add a default case for weekdays and print "It's a weekday."

	}
}
