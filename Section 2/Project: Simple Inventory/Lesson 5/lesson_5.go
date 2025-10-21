package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Product struct {
	Price    float64
	Quantity int
}

func generateReport(inventory map[string]Product, report string, threshold int) {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	productData := scanner.Text()
	productEntries := strings.Split(productData, ",")

	scanner.Scan()
	updateListOfProducts := scanner.Text()
	configurations := strings.Split(updateListOfProducts, ",")

	reportType := configurations[0]
	treshold, _ := strconv.Atoi(configurations[1])

	totalItems, inventoryValue := 0, 0.00

	inventory := make(map[string]Product)
	for _, product := range productEntries {
		prod := strings.Split(product, ":")
		name := prod[0]
		price, _ := strconv.ParseFloat(prod[1], 64)
		quantity, _ := strconv.Atoi(prod[2])
		inventory[name] = Product{price, quantity}
		totalItems += quantity
		inventoryValue += float64(quantity) * price
	}
	averageValue := inventoryValue / float64(totalItems)

	switch reportType {
	case "full":
		fmt.Println("=== FULL INVENTORY REPORT ===")
	case "low_stock":
		fmt.Println("=== LOW STOCK REPORT ===")
	case "high_value":
		fmt.Println("=== HIGH VALUE REPORT ===")
	}

	fmt.Printf("Total Products: %d\n", len(productEntries))
	fmt.Printf("Total Items in Stock: %d\n", totalItems)
	fmt.Printf("Total Inventory Value: $%.2f\n", inventoryValue)
	fmt.Printf("Average Product Price: $%.2f\n", averageValue)

	keys := make([]string, 0)
	switch reportType {
	case "full":
		fmt.Println("All Products:")
		for product := range inventory {
			keys = append(keys, product)
		}
	case "low_stock":
		fmt.Printf("Products with stock ≤ %d:\n", treshold)
		for product := range inventory {
			if inventory[product].Quantity <= treshold {
				keys = append(keys, product)
			}
		}
	case "high_value":
		fmt.Printf("Products with value ≥ $%.2f:\n", float64(treshold))
		for product := range inventory {
			if float64(inventory[product].Quantity)*inventory[product].Price >= float64(treshold) {
				keys = append(keys, product)
			}
		}
	}

	sort.Strings(keys)
	for _, product := range keys {
		fmt.Printf("- %s: $%.2f × %d = $%.2f\n", product, inventory[product].Price, inventory[product].Quantity, float64(inventory[product].Quantity)*inventory[product].Price)
	}

	quantity, price := 0, 0.00
	for _, prod := range keys {
		quantity += inventory[prod].Quantity
		price += float64(inventory[prod].Quantity) * inventory[prod].Price
	}

	fmt.Println("Filtered Results:")
	fmt.Printf("Products shown: %d\n", len(keys))
	fmt.Printf("Items in filtered products: %d\n", quantity)
	fmt.Printf("Value of filtered products: $%.2f\n", price)

	mostExp, leastExp, mostName, leastName := 0.00, math.MaxFloat64, "", ""
	highStock, lowStock, highName, lowName := 0, math.MaxInt, "", ""
	inventoryLowStock, inventoryHighStock, inventoryHighValue := 0, 0, 0
	for product := range inventory {
		if inventory[product].Price > mostExp {
			mostExp = inventory[product].Price
			mostName = product
		}
		if inventory[product].Price < leastExp {
			leastExp = inventory[product].Price
			leastName = product
		}
		if inventory[product].Quantity > highStock {
			highStock = inventory[product].Quantity
			highName = product
		}
		if inventory[product].Quantity < lowStock {
			lowStock = inventory[product].Quantity
			lowName = product
		}
		if inventory[product].Quantity <= 5 {
			inventoryLowStock += 1
		}
		if inventory[product].Quantity > 20 {
			inventoryHighStock += 1
		}
		if inventory[product].Price*float64(inventory[product].Quantity) >= 500.00 {
			inventoryHighValue += 1
		}
	}
	fmt.Println("Price Analysis:")
	fmt.Printf("Most expensive: %s at $%.2f\n", mostName, mostExp)
	fmt.Printf("Least expensive: %s at $%.2f\n", leastName, leastExp)

	fmt.Printf("Stock Analysis:\n")
	fmt.Printf("Highest stock: %s with %d units\n", highName, highStock)
	fmt.Printf("Lowest stock: %s with %d units\n", lowName, lowStock)

	fmt.Printf("Low stock items (≤5): %d\n", inventoryLowStock)
	fmt.Printf("High stock items (>20): %d\n", inventoryHighStock)
	fmt.Printf("High value items (≥$500): %d\n", inventoryHighValue)

	fmt.Printf("Report generated successfully\n")
	fmt.Printf("Threshold applied: %.2f\n", float64(treshold))
}
