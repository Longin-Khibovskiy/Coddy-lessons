package main

import "fmt"

// getGreeting returns a greeting based on the hour of the day
func getGreeting(hour int) string {
	// TODO: Return "Good morning" if hour is less than 12
	// TODO: Return "Good afternoon" if hour is between 12 and 17 (inclusive)
	// TODO: Return "Good evening" for all other hours
	if hour < 12 {
		return "Good morning"
	} else if hour >= 12 && hour <= 17 {
		return "Good afternoon"
	} else {
		return "Good evening"
	}
	// Your code here

}

func main() {
	// Test the function with different hours
	morningHour := 8
	afternoonHour := 15
	eveningHour := 20

	fmt.Println("At", morningHour, "hours:", getGreeting(morningHour))
	fmt.Println("At", afternoonHour, "hours:", getGreeting(afternoonHour))
	fmt.Println("At", eveningHour, "hours:", getGreeting(eveningHour))
}
