package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Read input
	var numOpsStr string
	var operationsStr string
	fmt.Scanln(&numOpsStr)
	fmt.Scanln(&operationsStr)

	var scores map[string]int
	var grades map[string]int

	operations := strings.Split(operationsStr, ",")

	for _, operation := range operations {
		operationData := strings.Split(operation, ":")
		oper := operationData[0]
		switch oper {
		case "check":
			switch operationData[1] {
			case "scores":
				if scores == nil {
					fmt.Println("Map scores is nil")
				} else if scores != nil {
					fmt.Println("Map scores is initialized")
				}
			case "grades":
				if grades == nil {
					fmt.Println("Map grades is nil")
				} else if grades != nil {
					fmt.Println("Map grades is initialized")
				}
			}
		case "init":
			switch operationData[1] {
			case "scores":
				scores = make(map[string]int)
				fmt.Println("Initialized map scores")
			case "grades":
				grades = make(map[string]int)
				fmt.Println("Initialized map grades")
			}
		case "add":
			switch operationData[1] {
			case "scores":
				if scores == nil {
					fmt.Println("Cannot add to nil map scores")
				} else {
					value, _ := strconv.Atoi(operationData[3])
					scores[operationData[2]] = value
					fmt.Printf("Added %s:%d to scores\n", operationData[2], value)
				}

			case "grades":
				if grades == nil {
					fmt.Println("Cannot add to nil map grades")
				} else {
					value, _ := strconv.Atoi(operationData[3])
					grades[operationData[2]] = value
					fmt.Printf("Added %s:%d to grades\n", operationData[2], value)
				}
			}

		}

	}
	if scores == nil {
		fmt.Println("Final state - scores: nil")
	} else if scores != nil {
		fmt.Printf("Final state - scores: %d entries\n", len(scores))
	}
	if grades == nil {
		fmt.Println("Final state - grades: nil")
	} else if grades != nil {
		fmt.Printf("Final state - grades: %d entries", len(grades))
	}
}
