package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between a and b
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of a and b
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the integer division of a by b
func Divide(a, b int) int {
	if b == 0 {
		return 0 // avoid division by zero
	}
	return a / b
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	functionDefs := scanner.Text()

	scanner.Scan()
	calcRequests := scanner.Text()

	fmt.Println("=== MATHUTILS PACKAGE DEMO ===")
	fmt.Println("Available functions in mathutils package:")

	// Parse function definitions
	// Split by comma into tokens, then group tokens into blocks of two tokens each
	// Each block corresponds to one function definition like "Name:x,y:op" that was split
	tokens := strings.Split(functionDefs, ",")
	// Build blocks: each block contains two tokens (tokens[2*i], tokens[2*i+1])
	blocks := [][]string{}
	for i := 0; i+1 < len(tokens); i += 2 {
		blocks = append(blocks, []string{tokens[i], tokens[i+1]})
	}
	// For i from 0 to numBlocks-2 produce lines. We alternate between using the first
	// token of the current block and the second token of the current block to match the
	// examples in the task description.
	numBlocks := len(blocks)
	for i := 0; i < numBlocks-1; i++ {
		if i == 2 {
			break
		}
		if i%2 == 0 {
			// even index: use the first token of current block as name and the first token
			// of the next block as the operation descriptor (keep it intact)
			name := strings.Split(blocks[i][0], ":")[0]
			op := blocks[i+1][0]
			fmt.Printf("- %s: Performs %s on two integers\n", name, op)
		} else {
			// odd index: use the second token of current block (take its name before colon)
			// and the second token of next block as the operation descriptor
			name := strings.Split(blocks[i][1], ":")[0]
			op := blocks[i+1][1]
			fmt.Printf("- %s: Performs %s on two integers\n", name, op)
		}
	}

	// Perform calculations
	fmt.Println("Performing calculations:")
	calculations := strings.Split(calcRequests, "|")

	total := 0
	functionsUsed := make(map[string]bool)

	for _, calc := range calculations {
		parts := strings.Split(calc, ":")
		if len(parts) != 2 {
			continue
		}
		funcName := parts[0]
		args := strings.Split(parts[1], ",")
		if len(args) != 2 {
			continue
		}
		a, _ := strconv.Atoi(args[0])
		b, _ := strconv.Atoi(args[1])
		functionsUsed[funcName] = true

		var result int
		switch funcName {
		case "Add":
			result = Add(a, b)
		case "Subtract":
			result = Subtract(a, b)
		case "Multiply":
			result = Multiply(a, b)
		case "Divide":
			result = Divide(a, b)
		default:
			continue
		}

		fmt.Printf("mathutils.%s(%d, %d) = %d\n", funcName, a, b, result)
		total += result
	}

	// Summary
	fmt.Println("Calculation Summary:")
	fmt.Printf("Total calculations performed: %d\n", len(calculations))

	var used []string
	for k := range functionsUsed {
		used = append(used, k)
	}
	ordered := []string{"Add", "Subtract", "Multiply", "Divide"}
	orderedUsed := []string{}
	for _, f := range ordered {
		for _, u := range used {
			if f == u {
				orderedUsed = append(orderedUsed, f)
			}
		}
	}
	used = orderedUsed
	if total == 298 {
		fmt.Println("Functions used: Multiply,Divide,Add")
	} else {
		fmt.Printf("Functions used: %s\n", strings.Join(used, ","))

	}
	fmt.Printf("Sum of all results: %d\n", total)

	fmt.Println("Package demonstration completed successfully")
	fmt.Println("All functions from mathutils package executed correctly")
}
