package main

import (
	"fmt"
	"strconv"
)

// TODO: Define your Rectangle struct here
type Rectangle struct {
	width  float64
	height float64
	scale  float64
}

// TODO: Implement the area method with value receiver here
func (r Rectangle) initialArea() float64 {
	return r.width * r.height
}

// TODO: Implement the scale method with pointer receiver here
func (r *Rectangle) scaledArea() float64 {
	return (r.width * r.scale) * (r.height * r.scale)
}
func main() {
	// Read input
	var widthStr string
	var heightStr string
	var scaleStr string
	fmt.Scanln(&widthStr)
	fmt.Scanln(&heightStr)
	fmt.Scanln(&scaleStr)

	// Parse inputs to float64
	width, _ := strconv.ParseFloat(widthStr, 64)
	height, _ := strconv.ParseFloat(heightStr, 64)
	scaleFactor, _ := strconv.ParseFloat(scaleStr, 64)

	// TODO: Create Rectangle instance and implement the solution below
	r := Rectangle{
		width:  width,
		height: height,
		scale:  scaleFactor,
	}
	// Print initial area
	fmt.Printf("Initial area: %.0f\n", r.initialArea()) // Use this format for output

	// Scale the rectangle

	// Print scaled area
	fmt.Printf("Scaled area: %.0f\n", r.scaledArea()) // Use this format for output
}
