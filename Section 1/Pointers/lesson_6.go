package main

import "fmt"

func main() {
	// This is a nil pointer to an integer
	var ptr *int

	// TODO: Check if ptr is nil
	// If it is nil, print "The pointer is nil"
	// If it is not nil, print "The pointer is not nil"
	if ptr == nil {
		fmt.Println("The pointer is nil")
	} else if ptr != nil {
		fmt.Println("The pointer is not nil")
	}
	fmt.Println("Program completed")
}
