package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Price    float64
	Quantity int
}

func checkStock(inventory map[string]Product, name string) (int, error) {
	if _, exists := inventory[name]; exists {
		return inventory[name].Quantity, nil
	}
	return 0, fmt.Errorf("product not found: %s", name)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	productData := scanner.Text()

	scanner.Scan()
	listOfProducts := scanner.Text()
	listProducts := strings.Split(listOfProducts, ",")

	productEntries := strings.Split(productData, ",")

	units := 0
	inventory := make(map[string]Product)
	for _, product := range productEntries {
		prod := strings.Split(product, ":")
		prodName := prod[0]
		prodPrice, _ := strconv.ParseFloat(prod[1], 64)
		prodQuantity, _ := strconv.Atoi(prod[2])
		inventory[prodName] = Product{prodPrice, prodQuantity}
	}

	productsFound, productsNotFound := 0, 0
	listNotFoundProducts := make([]string, 0, len(listProducts))
	fmt.Println("Stock Check Results:")
	for _, prod := range listProducts {
		_, err := checkStock(inventory, prod)
		if err != nil {
			fmt.Printf("%s: Error - %s\n", prod, err)
			productsNotFound += 1
			if prod != "" {
				listNotFoundProducts = append(listNotFoundProducts, prod)
			}
		} else {
			fmt.Printf("%s: %d units in stock\n", prod, inventory[prod].Quantity)
			units += inventory[prod].Quantity
			productsFound += 1
		}
	}

	fmt.Println("Check Summary:")
	fmt.Printf("Products checked: %d\n", len(listProducts))
	fmt.Printf("Products found: %d\n", productsFound)
	fmt.Printf("Products not found: %d\n", productsNotFound)

	fmt.Printf("Total stock for found products: %d units\n", units)
	if len(listNotFoundProducts) > 0 {
		fmt.Printf("Missing products: %s\n", strings.Join(listNotFoundProducts, ","))
	} else {
		fmt.Println("All requested products are available")
	}
}
