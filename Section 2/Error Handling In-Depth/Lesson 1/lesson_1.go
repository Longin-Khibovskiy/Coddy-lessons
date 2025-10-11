package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Read input
	var userData string
	var validationRules string
	fmt.Scanln(&userData)
	fmt.Scanln(&validationRules)

	// Parse user data
	userParts := strings.Split(userData, ",")
	username := userParts[0]
	email := userParts[1]
	ageStr := userParts[2]

	// Parse validation rules
	ruleParts := strings.Split(validationRules, ",")
	minUsernameLength, _ := strconv.Atoi(ruleParts[0])
	maxAge, _ := strconv.Atoi(ruleParts[1])

	// Convert age to integer
	age, _ := strconv.Atoi(ageStr)

	err := ValidateUser(username, email, age, minUsernameLength, maxAge)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Validation Summary:")
		fmt.Println("Username:", username)
		fmt.Println("Email:", email)
		fmt.Println("Age:", age)
		fmt.Println("Status: FAILED")
	} else {
		fmt.Println("User validation successful")
		fmt.Println("Validation Summary:")
		fmt.Println("Username:", username)
		fmt.Println("Email:", email)
		fmt.Println("Age:", age)
		fmt.Println("Status: PASSED")
	}
}

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("Validation failed for %s: %s", e.Field, e.Message)
}

func ValidateUser(username string, email string, age int, usernameLength int, maxAge int) error {
	if len(username) < usernameLength {
		return ValidationError{"username", "too short"}
	}
	if !strings.Contains(email, "@") {
		return ValidationError{"email", "must contain @ symbol"}
	}
	if age < 18 || age > maxAge {
		return ValidationError{"age", "must be between 18 and maximum age"}
	}
	return nil
}
