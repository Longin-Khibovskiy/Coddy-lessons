package main

import (
	"fmt"
)

// Counter is a simple counter data structure
type Counter struct {
	value int
	name  string
}

// NewCounter creates a new counter with a given name
func NewCounter(name string) Counter {
	return Counter{
		value: 0,
		name:  name,
	}
}

// Increment increases the counter by 1
func (c *Counter) Increment() {
	c.value++
}

// Decrement decreases the counter by 1 (but not below 0)
func (c *Counter) Decrement() {
	// Add code to decrement the counter's value
	// Make sure the value doesn't go below 0
	if c.value > 0 {
		c.value--
	}
}

// Reset sets the counter back to 0
func (c *Counter) Reset() {
	c.value = 0
}

// Value returns the current count
func (c Counter) Value() int {
	// Add code to return the current value
	return c.value // Replace this
}

// String returns a string representation of the counter
func (c Counter) String() string {
	return fmt.Sprintf("%s: %d", c.name, c.value)
}

func main() {
	// Create a new counter named "Visitors"
	visitors := NewCounter("Visitors")

	// Increment the counter a few times
	visitors.Increment()
	visitors.Increment()
	visitors.Increment()

	// Print the current count
	fmt.Println(visitors)
	fmt.Println("Current value:", visitors.Value())

	// Decrement the counter
	visitors.Decrement()
	fmt.Println(visitors)

	// Reset the counter
	visitors.Reset()
	fmt.Println(visitors)

	// Test that counter doesn't go below zero
	visitors.Decrement()
	fmt.Println(visitors)
}
