package main

import "fmt"

func main() {
	// A slice with 3 elements but capacity of 5
	numbers := make([]int, 3, 5)

	// TODO: Print the length of the slice using len()
	fmt.Println("Length:", len(numbers))

	// TODO: Print the capacity of the slice using cap()
	fmt.Println("Capacity:", cap(numbers))

}
