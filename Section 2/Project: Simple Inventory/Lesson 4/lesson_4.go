package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Product struct {
	Price    float64
	Quantity int
}

func updateStock(inventory map[string]Product, name string, quantity int) error {
	if _, exists := inventory[name]; !exists {
		return fmt.Errorf("product not found: %s", name)
	}
	if inventory[name].Quantity+quantity < 0 {
		return fmt.Errorf("insufficient stock: cannot reduce %s by %d, only %d available", name, -quantity, inventory[name].Quantity)
	}
	res := inventory[name].Quantity + quantity
	inventory[name] = Product{Price: inventory[name].Price, Quantity: res}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	productData := scanner.Text()
	productEntries := strings.Split(productData, ",")

	scanner.Scan()
	updateListOfProducts := scanner.Text()
	updateProducts := strings.Split(updateListOfProducts, ",")

	successfullUpdate := 0
	failedUpdate := 0
	itemsStock := 0
	valueStock := 0.00
	failedProcess := make([]string, 0)

	inventory := make(map[string]Product)
	for _, product := range productEntries {
		prod := strings.Split(product, ":")
		name := prod[0]
		price, _ := strconv.ParseFloat(prod[1], 64)
		quantity, _ := strconv.Atoi(prod[2])
		inventory[name] = Product{price, quantity}
	}

	fmt.Println("Processing Stock Updates:")
	for _, product := range updateProducts {
		prod := strings.Split(product, ":")
		name := prod[0]
		quantity, _ := strconv.Atoi(prod[1])
		err := updateStock(inventory, name, quantity)
		if err == nil {
			if quantity > 0 {
				fmt.Printf("%s: Added %d units - New stock: %d\n", name, quantity, inventory[name].Quantity)
			} else if quantity == 0 {
				fmt.Printf("%s: No change - Current stock: %d\n", name, inventory[name].Quantity)
			} else if quantity < 0 {
				fmt.Printf("%s: Removed %d units - New stock: %d\n", name, -quantity, inventory[name].Quantity)
			}
			successfullUpdate += 1
		} else {
			fmt.Printf("%s: Update failed - %s\n", name, err)
			failedUpdate += 1
			failedProcess = append(failedProcess, name)
		}
	}

	keys := make([]string, 0)
	for key := range inventory {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	fmt.Println("Update Summary:")
	fmt.Printf("Updates processed: %d\n", len(updateProducts))
	fmt.Printf("Updates successful: %d\n", successfullUpdate)
	fmt.Printf("Updates failed: %d\n", failedUpdate)

	fmt.Println("Updated Inventory:")
	for _, productName := range keys {
		fmt.Printf("- %s: $%.2f (Stock: %d)\n", productName, inventory[productName].Price, inventory[productName].Quantity)
		itemsStock += inventory[productName].Quantity
		valueStock += float64(inventory[productName].Quantity) * inventory[productName].Price
	}

	fmt.Println("Final Inventory Statistics:")
	fmt.Printf("Total Products: %d\n", len(inventory))
	fmt.Printf("Total Items in Stock: %d\n", itemsStock)
	fmt.Printf("Total Inventory Value: $%.2f\n", valueStock)
	if len(failedProcess) > 0 {
		fmt.Printf("Failed updates: %s\n", strings.Join(failedProcess, ","))
	} else {
		fmt.Println("All stock updates were processed successfully")
	}
}
