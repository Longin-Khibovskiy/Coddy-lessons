package main

import (
	"fmt"
	"strconv"
)

type Vehicle interface {
	GetInfo() string
}

// 2. Define the Car struct
type Car struct {
	Brand string
	Speed int
}

// 3. Implement the GetInfo() method for Car
func (c Car) GetInfo() string {
	return fmt.Sprintf("Brand: %s, Speed: %d km/h", c.Brand, c.Speed)
}
func main() {
	// Read input
	var brand string
	var speedStr string
	fmt.Scanln(&brand)
	fmt.Scanln(&speedStr)

	// Convert speed string to integer
	speed, _ := strconv.Atoi(speedStr)

	// TODO: Write your code below
	// 1. Define the Vehicle interface
	// 4. Create a Car instance and call GetInfo()
	var v Vehicle = Car{
		Brand: brand,
		Speed: speed,
	}
	result := v.GetInfo()
	// Output the result
	fmt.Println(result)
}
