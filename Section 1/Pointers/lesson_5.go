package main

import (
	"fmt"
	"strings"
)

// makeUppercase takes a pointer to a string and changes the string to uppercase
func makeUppercase(strPtr *string) {
	// TODO: Use strings.ToUpper() to change the string that strPtr points to
	// Hint: You need to dereference the pointer to access the string value
	*strPtr = strings.ToUpper(*strPtr)

}

func main() {
	// We already have a string variable
	message := "hello, world"

	// Print the original message
	fmt.Printf("Original message: %s\n", message)

	// Call makeUppercase with the address of message
	makeUppercase(&message)

	// Print the modified message
	fmt.Printf("Modified message: %s\n", message)
}
