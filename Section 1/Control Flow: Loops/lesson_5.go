package main

import "fmt"

func main() {
	// Loop through numbers 1 to 10
	for i := 1; i <= 10; i++ {
		// TODO: Use the continue keyword to skip numbers that are divisible by 3
		// Hint: Use the modulo operator (%) to check if i is divisible by 3
		if i%3 == 0 {
			continue
		}

		fmt.Println(i)
	}
}
