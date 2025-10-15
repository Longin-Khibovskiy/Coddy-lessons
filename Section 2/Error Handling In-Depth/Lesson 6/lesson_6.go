package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type DivisionError struct {
	Dividend float64
	Divisor  float64
}

type ValidationError struct {
	Field string
	Value float64
}

func (divisionError DivisionError) Error() string {
	return fmt.Sprintf("division error: cannot divide %.0f by %.0f", divisionError.Dividend, divisionError.Divisor)
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf("validation error: invalid %s value %.0f", validationError.Field, validationError.Value)
}

func safeDivide(dividend float64, divisor float64, mode string) (float64, error) {
	if divisor == 0 {
		return 0, DivisionError{dividend, divisor}
	}
	if dividend < 0 && mode == "strict" {
		return 0, ValidationError{"dividend", dividend}
	}
	if divisor < 0 && mode == "strict" {
		return 0, ValidationError{"divisor", divisor}
	}
	return dividend / divisor, nil
}

func performCalculation(dividend float64, divisor float64, calculatorName string, mode string) (float64, error) {
	result, err := safeDivide(dividend, divisor, mode)
	if err != nil {
		return 0, fmt.Errorf("calculation failed in %s: %w", calculatorName, err)
	}
	return result, nil
}

func main() {
	// Read input
	var operationInput string
	var settingsInput string
	fmt.Scanln(&operationInput)
	fmt.Scanln(&settingsInput)

	// Parse operation input (operation,num1,num2,precision)
	operationParts := strings.Split(operationInput, ",")
	operation := operationParts[0]
	num1, _ := strconv.ParseFloat(operationParts[1], 64)
	num2, _ := strconv.ParseFloat(operationParts[2], 64)
	precision, _ := strconv.Atoi(operationParts[3])

	settingsParts := strings.Split(settingsInput, ",")
	calculatorName := settingsParts[0]
	mode := settingsParts[1]

	res, err := performCalculation(num1, num2, calculatorName, mode)
	if err == nil {
		fmt.Printf("Calculation successful: %.*f\n", precision, res)
	} else {
		fmt.Printf("Calculation failed: %v\n", err)
	}

	var divisionError DivisionError
	var validationError ValidationError
	isDivisionError := errors.As(err, &divisionError)
	isValidationError := errors.As(err, &validationError)
	fmt.Println("Checking for division error:", isDivisionError)
	fmt.Println("Checking for validation error:", isValidationError)
	if isDivisionError {
		fmt.Println("Division Error Details:")
		fmt.Println("Dividend:", divisionError.Dividend)
		fmt.Println("Divisor:", divisionError.Divisor)
	}
	if isValidationError {
		fmt.Println("Validation Error Details:")
		fmt.Println("Field:", validationError.Field)
		fmt.Println("Value:", validationError.Value)
	}
	status := "Success"
	if err != nil {
		status = "Failed"
	}
	fmt.Printf("Calculator Summary:\n")
	fmt.Printf("Name: %s\n", calculatorName)
	fmt.Printf("Mode: %s\n", mode)
	fmt.Printf("Operation: %s\n", operation)
	fmt.Printf("Input: %g / %g\n", num1, num2)
	fmt.Printf("Precision: %d decimal places\n", precision)
	fmt.Printf("Status: %s\n", status)
}
