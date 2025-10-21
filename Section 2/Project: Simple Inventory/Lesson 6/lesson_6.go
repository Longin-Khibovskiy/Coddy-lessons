package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Product struct {
	Price    float64
	Quantity int
}

func checkStock(inventory map[string]Product, name string) (int, error) {
	if product, exists := inventory[name]; exists {
		return product.Quantity, nil
	}
	return 0, errors.New("product not found in inventory")
}

func addNewItem(inventory map[string]Product, param string) error {
	parts := strings.Split(param, ":")
	if len(parts) != 3 {
		return errors.New("invalid product format")
	}
	name := parts[0]
	price, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return errors.New("invalid price")
	}
	quantity, err := strconv.Atoi(parts[2])
	if err != nil {
		return errors.New("invalid quantity")
	}
	if _, exists := inventory[name]; exists {
		return errors.New("product already exists")
	}
	inventory[name] = Product{price, quantity}
	return nil
}

func updateStock(inventory map[string]Product, param string) error {
	parts := strings.Split(param, ":")
	if len(parts) != 2 {
		return errors.New("invalid update format")
	}
	name := parts[0]
	change, err := strconv.Atoi(parts[1])
	if err != nil {
		return errors.New("invalid stock change")
	}
	product, exists := inventory[name]
	if !exists {
		return errors.New("product not found")
	}
	newQuantity := product.Quantity + change
	if newQuantity < 0 {
		return errors.New("insufficient stock to remove")
	}
	product.Quantity = newQuantity
	inventory[name] = product
	return nil
}

func generateReport(inventory map[string]Product, param string) {
	parts := strings.Split(param, ",")
	if len(parts) < 2 {
		fmt.Println("Invalid report parameters")
		return
	}
	reportType := strings.ToLower(strings.TrimSpace(parts[0]))
	threshold, _ := strconv.Atoi(parts[1])

	// normalize synonyms
	switch reportType {
	case "low":
		reportType = "low"
	case "high":
		reportType = "high_value"
	}

	fmt.Printf("Generating %s report with threshold %d\n", reportType, threshold)

	keys := make([]string, 0)
	for k := range inventory {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	switch reportType {
	case "low":
		fmt.Println("=== LOW REPORT ===")
		fmt.Printf("Products with stock below %d:\n", threshold)
		for _, k := range keys {
			p := inventory[k]
			if p.Quantity < threshold {
				fmt.Printf("- %s: %d units (Price: $%.2f)\n", k, p.Quantity, p.Price)
			}
		}
	case "full":
		fmt.Println("=== FULL REPORT ===")
		fmt.Printf("Complete inventory (minimum threshold: %d):\n", threshold)
		for _, k := range keys {
			p := inventory[k]
			status := "[OK]"
			if p.Quantity < threshold {
				status = "[LOW STOCK]"
			}
			fmt.Printf("- %s: %d units @ $%.2f each %s\n", k, p.Quantity, p.Price, status)
		}
	case "high_value":
		fmt.Println("=== HIGH VALUE REPORT ===")
		fmt.Printf("Products with total value ≥ $%d:\n", threshold)
		for _, k := range keys {
			p := inventory[k]
			value := p.Price * float64(p.Quantity)
			if value >= float64(threshold) {
				fmt.Printf("- %s: $%.2f ($%.2f × %d)\n", k, value, p.Price, p.Quantity)
			}
		}
	default:
		fmt.Println("Unknown report type")
	}
}

func main() {
	var input1, input2, input3 string
	fmt.Scanln(&input1)
	fmt.Scanln(&input2)
	fmt.Scanln(&input3)

	inventory := make(map[string]Product)
	products := strings.Split(input1, ",")
	for _, item := range products {
		parts := strings.Split(item, ":")
		name := parts[0]
		price, _ := strconv.ParseFloat(parts[1], 64)
		quantity, _ := strconv.Atoi(parts[2])
		inventory[name] = Product{price, quantity}
	}

	fmt.Println("=== INVENTORY MANAGEMENT SYSTEM ===")
	fmt.Printf("System initialized with %d products\n", len(inventory))
	fmt.Println("Starting interactive session...")

	operations := strings.Split(input2, ",")
	params := strings.Split(input3, "|")

	for i, op := range operations {
		param := ""
		if i < len(params) {
			param = params[i]
		}

		switch op {
		case "check":
			fmt.Println("--- STOCK CHECK ---")
			fmt.Printf("Checking stock for: %s\n", param)
			qty, err := checkStock(inventory, param)
			if err != nil {
				fmt.Println("Product not found in inventory")
			} else {
				fmt.Printf("Stock level: %d units\n", qty)
			}
		case "add":
			fmt.Println("--- ADD ITEM ---")
			parts := strings.Split(param, ":")
			fmt.Printf("Adding new product: %s\n", parts[0])
			err := addNewItem(inventory, param)
			if err != nil {
				fmt.Printf("Failed to add product: %s\n", err.Error())
			} else {
				fmt.Println("Product added successfully")
			}
		case "update":
			fmt.Println("--- UPDATE STOCK ---")
			parts := strings.Split(param, ":")
			fmt.Printf("Updating stock for: %s\n", parts[0])
			err := updateStock(inventory, param)
			if err != nil {
				fmt.Printf("Update failed: %s\n", err.Error())
			} else {
				change, _ := strconv.Atoi(parts[1])
				newQty := inventory[parts[0]].Quantity
				if change >= 0 {
					fmt.Printf("Added %d units. New stock: %d\n", change, newQty)
				} else {
					fmt.Printf("Removed %d units. New stock: %d\n", -change, newQty)
				}
			}
		case "report":
			fmt.Println("--- GENERATE REPORT ---")
			generateReport(inventory, param)
		case "exit":
			fmt.Println("--- SYSTEM EXIT ---")
			totalProducts := len(inventory)
			totalItems, totalValue := 0, 0.0
			for _, p := range inventory {
				totalItems += p.Quantity
				totalValue += p.Price * float64(p.Quantity)
			}
			fmt.Println("Final inventory status:")
			fmt.Printf("Total products: %d\n", totalProducts)
			fmt.Printf("Total items: %d\n", totalItems)
			fmt.Printf("Total value: $%.2f\n", totalValue)
			fmt.Println("Session completed successfully")
			fmt.Println("Thank you for using the Inventory Management System")
			return
		default:
			fmt.Printf("Unknown operation: %s\n", op)
		}

		if op != "exit" {
			fmt.Println("Operation completed. Continuing to next operation...")
		}
	}
}
