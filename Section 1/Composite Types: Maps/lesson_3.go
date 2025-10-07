package main

import "fmt"

func main() {
	// TODO: Create a map called 'favoriteColors' using make
	// The keys should be strings (names) and the values should also be strings (colors)
	favoriteColors := make(map[string]string)
	// Adding some key-value pairs to the map
	favoriteColors["Alice"] = "Blue"
	favoriteColors["Bob"] = "Green"
	favoriteColors["Charlie"] = "Red"

	// Print the map
	fmt.Println("Favorite colors:", favoriteColors)

	// Print Bob's favorite color
	fmt.Println("Bob's favorite color is", favoriteColors["Bob"])
}
