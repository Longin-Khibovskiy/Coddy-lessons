package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numProductsStr string
	var productDataStr string
	var operationsStr string

	fmt.Scanln(&numProductsStr)
	fmt.Scanln(&productDataStr)
	fmt.Scanln(&operationsStr)

	type Product struct {
		Price    float64
		Quantity int
	}

	products := make(map[string]*Product)
	productNames := []string{}

	productsData := strings.Split(productDataStr, ",")

	for _, prod := range productsData {
		p := strings.Split(prod, ":")
		productName := p[0]
		productPrice, _ := strconv.ParseFloat(p[1], 64)
		productQuantity, _ := strconv.Atoi(p[2])

		products[productName] = &Product{productPrice, productQuantity}
		productNames = append(productNames, productName)
	}

	fmt.Println("Initial Inventory:")
	for _, product := range productNames {
		fmt.Printf("%s: $%.2f (Stock: %d)\n", product, products[product].Price, products[product].Quantity)
	}

	operationsData := strings.Split(operationsStr, ",")
	for _, operations := range operationsData {
		operation := strings.Split(operations, ":")
		operType := operation[0]
		operName := operation[1]
		switch operType {
		case "quantity":
			operValue, _ := strconv.Atoi(operation[2])
			products[operName].Quantity = operValue
			fmt.Printf("Updated %s: %s changed to %d\n", operName, operType, operValue)
		case "price":
			operValue, _ := strconv.ParseFloat(operation[2], 64)
			products[operName].Price = operValue
			fmt.Printf("Updated %s: %s changed to %.2f\n", operName, operType, operValue)
		}
	}

	fmt.Println("Updated Inventory:")
	totalValue := 0.00
	for _, product := range productNames {
		fmt.Printf("%s: $%.2f (Stock: %d)\n", product, products[product].Price, products[product].Quantity)
		totalValue += products[product].Price * float64(products[product].Quantity)
	}

	fmt.Printf("Total Inventory Value: $%.2f\n", totalValue)
}
