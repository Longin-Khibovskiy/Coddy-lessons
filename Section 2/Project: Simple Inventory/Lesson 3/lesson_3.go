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

func addNewItem(inventory map[string]Product, name string, price float64, quantity int) error {
	if _, exists := inventory[name]; exists {
		return fmt.Errorf("product already exists: %s", name)
	}
	inventory[name] = Product{price, quantity}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	productData := scanner.Text()
	productEntries := strings.Split(productData, ",")

	scanner.Scan()
	addListOfProducts := scanner.Text()
	addProducts := strings.Split(addListOfProducts, ",")

	addedItems := 0
	rejectedItems := 0
	totalQuantity := 0
	totalPrice := 0.00
	var rejectedProducts []string
	inventory := make(map[string]Product)
	for _, product := range productEntries {
		prod := strings.Split(product, ":")
		productName := prod[0]
		productPrice, _ := strconv.ParseFloat(prod[1], 64)
		productQuantity, _ := strconv.Atoi(prod[2])
		inventory[productName] = Product{productPrice, productQuantity}
	}

	fmt.Println("Adding New Items:")
	for _, newProduct := range addProducts {
		newProd := strings.Split(newProduct, ":")
		prodName := newProd[0]
		prodPrice, _ := strconv.ParseFloat(newProd[1], 64)
		prodQuantity, _ := strconv.Atoi(newProd[2])
		err := addNewItem(inventory, prodName, prodPrice, prodQuantity)
		if err != nil {
			fmt.Printf("%s: Failed to add - %s\n", prodName, err)
			rejectedItems += 1
			rejectedProducts = append(rejectedProducts, prodName)
		} else {
			fmt.Printf("%s: Successfully added - $%.2f (Stock: %d)\n", prodName, prodPrice, prodQuantity)
			addedItems += 1
		}
	}

	fmt.Println("Addition Summary:")
	fmt.Printf("Items processed: %d\n", len(addProducts))
	fmt.Printf("Items added: %d\n", addedItems)
	fmt.Printf("Items rejected: %d\n", rejectedItems)

	var keys []string
	for key := range inventory {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	fmt.Println("Updated Inventory:")
	for _, prod := range keys {
		fmt.Printf("- %s: $%.2f (Stock: %d)\n", prod, inventory[prod].Price, inventory[prod].Quantity)
		floatQuantity := float64(inventory[prod].Quantity)
		totalPrice += inventory[prod].Price * floatQuantity
		totalQuantity += inventory[prod].Quantity
	}

	fmt.Printf("Final Inventory Statistics:\n")
	fmt.Printf("Total Products: %d\n", len(keys))
	fmt.Printf("Total Items in Stock: %d\n", totalQuantity)
	fmt.Printf("Total Inventory Value: $%.2f\n", totalPrice)

	if len(rejectedProducts) > 0 {
		fmt.Printf("Rejected products: %s\n", strings.Join(rejectedProducts, ","))
	} else {
		fmt.Println("All new products were added successfully")
	}
}
