package main

import (
	"fmt"
	"strconv"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func printShapeInfo(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	// Read input
	var shapeType string
	var dimension1 string
	var dimension2 string
	fmt.Scanln(&shapeType)
	fmt.Scanln(&dimension1)
	fmt.Scanln(&dimension2)

	// Convert string inputs to float64
	dim1, _ := strconv.ParseFloat(dimension1, 64)
	dim2, _ := strconv.ParseFloat(dimension2, 64)

	// TODO: Write your code below
	// Define Shape interface, Circle and Rectangle structs
	// Implement Area() methods for both structs
	// Create printShapeInfo function
	// Create appropriate shape instance based on shapeType
	// Call printShapeInfo with the created shape
	var s Shape
	if shapeType == "circle" {
		s = Circle{dim1}
	} else if shapeType == "rectangle" {
		s = Rectangle{dim1, dim2}
	}

	printShapeInfo(s)
}
