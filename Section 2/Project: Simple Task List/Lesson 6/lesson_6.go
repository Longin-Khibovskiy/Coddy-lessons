package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

// checkStock returns the quantity of the product with the given name.
func checkStock(inventory map[string]Product, name string) (int, error) {
	product, ok := inventory[name]
	if !ok {
		return 0, errors.New("product not found in inventory")
	}
	return product.Quantity, nil
}

// addNewItem adds a new product if it does not exist.
func addNewItem(inventory map[string]Product, name string, price float64, quantity int) error {
	if _, exists := inventory[name]; exists {
		return errors.New("product already exists")
	}
	inventory[name] = Product{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
	return nil
}

// updateStock updates the stock by a change amount (positive or negative).
func updateStock(inventory map[string]Product, name string, change int) error {
	product, ok := inventory[name]
	if !ok {
		return errors.New("product not found in inventory")
	}
	if product.Quantity+change < 0 {
		return errors.New("insufficient stock to decrease")
	}
	product.Quantity += change
	inventory[name] = product
	return nil
}

// generateReport generates a report. If reportType is "below", lists products below threshold.
// If "above", lists products above or equal to threshold.
// If "full", lists all products with status based on threshold.
// If "low", lists products below threshold with detailed format.
func generateReport(inventory map[string]Product, reportType string, threshold int) ([]string, error) {
	var lines []string
	var keys []string
	for k := range inventory {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	switch reportType {
	case "below":
		for _, k := range keys {
			p := inventory[k]
			if p.Quantity < threshold {
				lines = append(lines, fmt.Sprintf("%s: %d units", p.Name, p.Quantity))
			}
		}
	case "above":
		for _, k := range keys {
			p := inventory[k]
			if p.Quantity >= threshold {
				lines = append(lines, fmt.Sprintf("%s: %d units", p.Name, p.Quantity))
			}
		}
	case "full":
		lines = append(lines, "=== FULL REPORT ===")
		lines = append(lines, fmt.Sprintf("Complete inventory (minimum threshold: %d):", threshold))
		for _, k := range keys {
			p := inventory[k]
			status := "[LOW STOCK]"
			if p.Quantity >= threshold {
				status = "[OK]"
			}
			lines = append(lines, fmt.Sprintf("- %s: %d units @ $%.2f each %s", p.Name, p.Quantity, p.Price, status))
		}
	case "low":
		lines = append(lines, "=== LOW REPORT ===")
		lines = append(lines, fmt.Sprintf("Products with stock below %d:", threshold))
		for _, k := range keys {
			p := inventory[k]
			if p.Quantity < threshold {
				lines = append(lines, fmt.Sprintf("- %s: %d units (Price: $%.2f)", p.Name, p.Quantity, p.Price))
			}
		}
	default:
		return nil, errors.New("invalid report type")
	}
	return lines, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input1 := scanner.Text()
	input1slice := strings.Split(input1, ",")
	inventory := make(map[string]Product)
	fmt.Printf("=== INVENTORY MANAGEMENT SYSTEM ===\nSystem initialized with %d products\nStarting interactive session...\n", len(input1slice))
	for _, value := range input1slice {
		valueslice := strings.Split(value, ":")
		price, _ := strconv.ParseFloat(valueslice[1], 64)
		quantity, _ := strconv.Atoi(valueslice[2])
		inventory[valueslice[0]] = Product{
			valueslice[0],
			price,
			quantity,
		}
	}
	scanner.Scan()
	input2 := scanner.Text()
	input2slice := strings.Split(input2, ",")

	scanner.Scan()
	input3 := scanner.Text()
	input3slice := strings.Split(input3, "|")

	for i, value := range input3slice {
		switch input2slice[i] {
		case "check":
			fmt.Println("--- STOCK CHECK ---")
			fmt.Printf("Checking stock for: %s\n", value)
			qty, err := checkStock(inventory, value)
			if err != nil {
				fmt.Println("Product not found in inventory")
			} else {
				fmt.Printf("Stock level: %d units\n", qty)
			}
			fmt.Println("Operation completed. Continuing to next operation...")
		case "add":
			fmt.Println("--- ADD ITEM ---")
			// value format: name:price:quantity
			parts := strings.Split(value, ":")
			if len(parts) != 3 {
				fmt.Println("Invalid add parameters")
				fmt.Println("Operation completed. Continuing to next operation...")
				continue
			}
			name := parts[0]
			price, err1 := strconv.ParseFloat(parts[1], 64)
			quantity, err2 := strconv.Atoi(parts[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid price or quantity")
				fmt.Println("Operation completed. Continuing to next operation...")
				continue
			}
			fmt.Printf("Adding new product: %s\n", name)
			err := addNewItem(inventory, name, price, quantity)
			if err != nil {
				fmt.Println("Failed to add item:", err)
			} else {
				fmt.Println("Product added successfully")
			}
			fmt.Println("Operation completed. Continuing to next operation...")
		case "update":
			fmt.Println("--- UPDATE STOCK ---")
			// value format: name:change
			parts := strings.Split(value, ":")
			if len(parts) != 2 {
				fmt.Println("Invalid update parameters")
				fmt.Println("Operation completed. Continuing to next operation...")
				continue
			}
			name := parts[0]
			change, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid change value")
				fmt.Println("Operation completed. Continuing to next operation...")
				continue
			}
			fmt.Printf("Updating stock for: %s\n", name)
			err = updateStock(inventory, name, change)
			if err != nil {
				fmt.Println("Failed to update stock:", err)
			} else {
				newQty := inventory[name].Quantity
				if change > 0 {
					fmt.Printf("Added %d units. New stock: %d\n", change, newQty)
				} else if change < 0 {
					fmt.Printf("Removed %d units. New stock: %d\n", -change, newQty)
				} else {
					fmt.Printf("No change in stock for %s. Current stock: %d\n", name, newQty)
				}
			}
			fmt.Println("Operation completed. Continuing to next operation...")
		case "report":
			fmt.Println("--- GENERATE REPORT ---")
			// value format: type,threshold
			parts := strings.Split(value, ",")
			if len(parts) != 2 {
				fmt.Println("Invalid report parameters")
				fmt.Println("Operation completed. Continuing to next operation...")
				continue
			}
			reportType := parts[0]
			threshold, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid threshold")
				fmt.Println("Operation completed. Continuing to next operation...")
				continue
			}
			fmt.Printf("Generating %s report with threshold %d\n", reportType, threshold)
			lines, err := generateReport(inventory, reportType, threshold)
			if err != nil {
				fmt.Println("Failed to generate report:", err)
			} else {
				if len(lines) == 0 {
					fmt.Println("No products match the report criteria.")
				} else {
					for _, line := range lines {
						fmt.Println(line)
					}
				}
			}
			fmt.Println("Operation completed. Continuing to next operation...")
		case "exit":
			fmt.Println("--- SYSTEM EXIT ---")
			// Compute stats
			totalProducts := len(inventory)
			totalItems := 0
			totalValue := 0.0
			for _, p := range inventory {
				totalItems += p.Quantity
				totalValue += float64(p.Quantity) * p.Price
			}
			fmt.Println("Final inventory status:")
			fmt.Printf("Total products: %d\n", totalProducts)
			fmt.Printf("Total items: %d\n", totalItems)
			fmt.Printf("Total value: $%.2f\n", totalValue)
			fmt.Println("Session completed successfully")
			fmt.Println("Thank you for using the Inventory Management System")
			return
		default:
			fmt.Println("Unknown operation:", input2slice[i])
			fmt.Println("Operation completed. Continuing to next operation...")
		}
	}
}
