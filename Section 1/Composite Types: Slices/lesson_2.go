package main

import "fmt"

func main() {
	// TODO: Create a slice literal named 'colors' containing the strings:
	// "red", "blue", "green", and "yellow"
	colors := [...]string{"red", "blue", "green", "yellow"}

	// Print the slice
	fmt.Println("Colors:", colors)

	// Print the length of the slice
	fmt.Println("Number of colors:", len(colors))
}
