package main

import (
	"fmt"
	"strconv"
)

func processNotification(data interface{}) {
	switch v := data.(type) {
	case string:
		fmt.Printf("Email notification: %s\n", v)
	case int:
		fmt.Printf("SMS notification with %d characters\n", v)
	case bool:
		fmt.Printf("Push notifications: %t\n", v)
	case float64:
		fmt.Println("Alert with priority:", v)
	default:
		fmt.Println("Unknown notification type")
	}
}
func main() {
	// Read input
	var notificationType string
	var content string
	fmt.Scanln(&notificationType)
	fmt.Scanln(&content)

	switch notificationType {
	case "email":
		processNotification(content)
	case "sms":
		value, _ := strconv.Atoi(content)
		processNotification(value)
	case "push":
		value, _ := strconv.ParseBool(content)
		if content == "enabled" {
			value = true
		}
		processNotification(value)
	case "alert":
		value, _ := strconv.ParseFloat(content, 64)
		processNotification(value)
	default:
		fmt.Println("Unknown notification type")
	}
	// TODO: Write your code below
	// 1. Create processNotification function that takes interface{} parameter
	// 2. Convert content to appropriate type based on notificationType
	// 3. Call processNotification with the converted value
	// 4. Use type switch inside processNotification to handle different types

}
