package main

import "fmt"

func main() {
	// Given variables
	age := 25
	score := 85
	isPremium := true

	// TODO: Create a boolean variable named 'isAdult' that checks if age is greater than or equal to 18
	isAdult := age >= 18
	// TODO: Create a boolean variable named 'isPassing' that checks if score is greater than or equal to 70
	isPassing := score >= 70
	// TODO: Create a boolean variable named 'isEligible' that checks if the person is both an adult AND has premium status
	isEligible := isAdult && isPremium
	// TODO: Create a boolean variable named 'needsImprovement' that checks if the score is less than 70 OR age is less than 20
	needsImprovement := score < 70 || age < 20
	// TODO: Create a boolean variable named 'isRejected' that is the NEGATION (opposite) of isPassing
	isRejected := !isPassing
	// Print results
	fmt.Println("Is Adult:", isAdult)
	fmt.Println("Is Passing:", isPassing)
	fmt.Println("Is Eligible:", isEligible)
	fmt.Println("Needs Improvement:", needsImprovement)
	fmt.Println("Is Rejected:", isRejected)
}
