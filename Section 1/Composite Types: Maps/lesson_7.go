package main

import "fmt"

func main() {
	// Map of inventory items and their quantities
	inventory := map[string]int{
		"pen":    15,
		"pencil": 10,
		"paper":  25,
		"eraser": 5,
	}

	fmt.Println("Initial inventory:", inventory)

	// TODO: Delete the 'pencil' entry from the inventory map
	delete(inventory, "pencil")

	fmt.Println("Updated inventory:", inventory)

	// Check if 'pencil' still exists in the map
	_, exists := inventory["pencil"]
	fmt.Println("Does 'pencil' exist in inventory?", exists)
}
