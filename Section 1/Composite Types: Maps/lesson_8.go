package main

import "fmt"

func main() {
	// A map of countries and their capitals
	countries := map[string]string{
		"France": "Paris",
		"Japan":  "Tokyo",
		"Brazil": "Brasília",
		"Canada": "Ottawa",
		"Egypt":  "Cairo",
	}

	// TODO: Use the len() function to get the number of countries in the map
	// and store it in a variable called countryCount
	countryCount := len(countries)

	// Print the number of countries
	fmt.Printf("The map contains %d countries\n", countryCount)
}
