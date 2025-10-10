package main

import (
	"fmt"
	"strconv"
	"strings"
)

func processData(anything interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", anything, anything)
}
func main() {
	// Read input
	var dataType string
	var valueStr string
	fmt.Scanln(&dataType)
	fmt.Scanln(&valueStr)

	// TODO: Write your code below
	// 1. Create the processData function that accepts interface{}
	// 2. Convert valueStr to the appropriate type based on dataType
	// 3. Call processData with the converted value
	switch dataType {
	case "string":
		processData(valueStr)
	case "int":
		val, _ := strconv.Atoi(valueStr)
		processData(val)
	case "float":
		val, _ := strconv.ParseFloat(valueStr, 64)
		processData(val)
	case "bool":
		val, _ := strconv.ParseBool(valueStr)
		processData(val)
	case "slice":
		value := strings.Split(valueStr, ",")
		var nums []int
		for _, num := range value {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}
		processData(nums)
	}
}
