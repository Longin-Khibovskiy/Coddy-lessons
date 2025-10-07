package main

import "fmt"

func main() {
	// TODO: Create a map literal called 'capitals' that maps countries to their capital cities
	// Include at least these three pairs:
	// "France" -> "Paris"
	// "Japan" -> "Tokyo"
	// "Brazil" -> "Brasilia"
	capitals := map[string]string{
		"France": "Paris",
		"Japan":  "Tokyo",
		"Brazil": "Brasilia",
	}
	// This will print out the map
	fmt.Println(capitals)

	// This will print out the capital of Japan
	fmt.Println("The capital of Japan is:", capitals["Japan"])
}
