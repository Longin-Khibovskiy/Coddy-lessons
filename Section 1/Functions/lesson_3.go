package main

import "fmt"

// This function prints a hello message
func sayHello() {
	fmt.Println("Hello, friend!")
}

// This function prints a goodbye message
func sayGoodbye() {
	fmt.Println("Goodbye, friend!")
}

// This function prints a thank you message
func sayThankYou() {
	fmt.Println("Thank you, friend!")
}

func main() {
	sayHello()
	sayThankYou()
	sayGoodbye()
}
