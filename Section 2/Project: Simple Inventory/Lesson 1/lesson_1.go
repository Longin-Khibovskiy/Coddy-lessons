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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	storeInfo := scanner.Text()

	scanner.Scan()
	productData := scanner.Text()

	storeData := strings.Split(storeInfo, ",")
	storeName := storeData[0]
	storeLocation := storeData[1]

	totalQuantity := 0
	totalValue := 0.00

	inventory := make(map[string]Product)
	for _, product := range strings.Split(productData, ",") {
		prod := strings.Split(product, ":")
		prodName := prod[0]
		prodPrice, _ := strconv.ParseFloat(prod[1], 64)
		prodQuantity, _ := strconv.Atoi(prod[2])
		inventory[prodName] = Product{prodPrice, prodQuantity}

		prodQuantityFloat, _ := strconv.ParseFloat(prod[2], 64)

		totalQuantity += prodQuantity
		totalValue += prodQuantityFloat * prodPrice
	}
	fmt.Printf("=== %s Inventory System ===\n", storeName)
	fmt.Printf("Location: %s\n", storeLocation)
	fmt.Printf("Inventory initialized with %d products\n", len(inventory))

	productsName := make([]string, 0)
	for key := range inventory {
		productsName = append(productsName, key)
	}
	sort.Strings(productsName)

	fmt.Println("Current Inventory:")
	for _, prod := range productsName {
		fmt.Printf("- %s: $%.2f (Stock: %d)\n", prod, inventory[prod].Price, inventory[prod].Quantity)
	}

	fmt.Println("Inventory Statistics:")
	fmt.Println("Total Products:", len(inventory))
	fmt.Printf("Total Items in Stock: %d\n", totalQuantity)
	fmt.Printf("Total Inventory Value: $%.2f\n", totalValue)
	fmt.Println("System Status: Ready\nInventory management system initialized successfully")
}
