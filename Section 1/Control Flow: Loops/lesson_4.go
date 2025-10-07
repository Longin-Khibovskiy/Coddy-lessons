package main

import "fmt"

func main() {
	// Loop through numbers 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Printf("Checking number: %d\n", i)

		// TODO: Check if the current number is divisible by 3
		// If it is, print "Found it: [number]!" and use break to exit the loop
		if i == 3 {
			fmt.Printf("Found it: %d!\n", i)
			break
		}
	}

	fmt.Println("Search complete")
}
