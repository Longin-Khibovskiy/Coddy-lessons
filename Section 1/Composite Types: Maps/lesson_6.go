package main

import "fmt"

func main() {
	// Map of student grades
	grades := map[string]int{
		"John":  85,
		"Sarah": 92,
		"Mike":  78,
		"Lisa":  95,
	}

	// Student to check
	studentToCheck := "Emma"

	// TODO: Check if studentToCheck exists in the grades map
	// and store the result in a variable called 'exists'
	// Hint: Use the comma-ok idiom
	_, exists := grades[studentToCheck]

	// Print the result
	if exists {
		fmt.Printf("%s's grade exists in the map\n", studentToCheck)
	} else {
		fmt.Printf("%s's grade does not exist in the map\n", studentToCheck)
	}
}
