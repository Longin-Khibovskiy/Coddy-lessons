package main

import (
	"fmt"
	"strconv"
)

// TODO: Define your Sensor struct here
type Sensor struct {
	ID          string
	Temperature float64
}

// TODO: Define your displayReading method with value receiver here
func (s Sensor) displayReading() {
	fmt.Printf("Sensor %s: %.1f°C\n", s.ID, s.Temperature)
}

// TODO: Define your adjustTemperature method with value receiver here
func (s Sensor) adjustTemperature(adjustment float64) {
	fmt.Printf("Adjusted reading: %.1f°C\n", s.Temperature+adjustment)
}
func main() {
	// Read inputs
	var sensorID string
	var tempStr string
	var adjustStr string

	fmt.Scanln(&sensorID)
	fmt.Scanln(&tempStr)
	fmt.Scanln(&adjustStr)

	// Parse temperature and adjustment values
	temperature, _ := strconv.ParseFloat(tempStr, 64)
	adjustment, _ := strconv.ParseFloat(adjustStr, 64)

	// TODO: Create a Sensor instance and call the required methods
	s := Sensor{
		ID:          sensorID,
		Temperature: temperature,
	}
	s.displayReading()
	s.adjustTemperature(adjustment)
	s.displayReading()
}
