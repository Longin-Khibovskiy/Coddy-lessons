package main

import "fmt"

func main() {
	// Candidate qualifications
	hasExperience := true
	hasEducation := true

	// TODO: Use the logical AND (&&) operator to check if the candidate
	// has both experience AND education, store the result in isQualified
	isQualified := hasEducation && hasExperience // Replace this line with your code

	// Print the result
	fmt.Println("Does the candidate qualify for the job?", isQualified)
}
