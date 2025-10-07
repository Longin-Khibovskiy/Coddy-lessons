package main

import (
	"fmt"
	"sort"
)

func main() {
	// Map of people and their favorite fruits
	favoriteFruits := map[string]string{
		"Alice": "Apple",
		"Bob":   "Banana",
		"Carol": "Cherry",
	}

	// TODO: Add a new entry for "David" with favorite fruit "Dragon Fruit"
	favoriteFruits["David"] = "Dragon Fruit"
	// TODO: Update Bob's favorite fruit to "Blueberry"
	favoriteFruits["Bob"] = "Blueberry"
	// Collect the names (keys) into a slice to sort them
	var names []string
	for person := range favoriteFruits {
		names = append(names, person)
	}
	sort.Strings(names)

	// Print the updated map in sorted order
	fmt.Println("Updated favorite fruits:")
	for _, person := range names {
		fruit := favoriteFruits[person]
		fmt.Printf("%s loves %s\n", person, fruit)
	}
}
