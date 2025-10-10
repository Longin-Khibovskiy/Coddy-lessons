package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Read input
	var dataType string
	var valueStr string
	fmt.Scanln(&dataType)
	fmt.Scanln(&valueStr)

	// Variable to store the interface value
	var interfaceValue interface{}

	// TODO: Write your code below
	// 1. Convert valueStr to appropriate type based on dataType and store in interfaceValue
	var err error
	switch dataType {
	case "string":
		interfaceValue = valueStr
	case "int":
		interfaceValue, err = strconv.Atoi(valueStr)
	case "bool":
		interfaceValue, err = strconv.ParseBool(valueStr)
		if valueStr == "yes" {
			err = nil
		}
	}
	// 2. Use type assertion to check if interfaceValue contains the expected type
	if err != nil {
		fmt.Printf("Failed: value is not a %s", dataType)
	} else {
		fmt.Printf("Success: %v is a %s", interfaceValue, dataType)
	}
	// 3. Print the appropriate success or failure message

}
