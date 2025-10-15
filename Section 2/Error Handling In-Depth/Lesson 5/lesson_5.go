package main

import (
	"fmt"
	"strconv"
	"strings"
)

func handleRequest(endpoint string, method string, action string, serverName string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Server %s recovered from panic: %s\n", serverName, r)
		}
	}()

	switch action {
	case "process":
		fmt.Printf("Processing %s request to %s\n", method, endpoint)
	case "panic_nil":
		fmt.Println("Attempting dangerous operation")
		panic("nil pointer access")
	case "panic_overflow":
		fmt.Println("Attempting resource allocation")
		panic("memory overflow")
	case "panic_timeout":
		fmt.Println("Attempting database connection")
		panic("connection timeout")
	default:
		fmt.Println("Unknown action:", action)
		panic("invalid action")
	}
}

func main() {
	var requestDetails string
	var serverConfig string
	fmt.Scanln(&requestDetails)
	fmt.Scanln(&serverConfig)

	requestParts := strings.Split(requestDetails, ",")
	endpoint := requestParts[0]
	method := requestParts[1]
	action := requestParts[2]

	configParts := strings.Split(serverConfig, ",")
	serverName := configParts[0]
	maxRequests, _ := strconv.Atoi(configParts[1])

	fmt.Printf("Starting %s with max %d requests\n", serverName, maxRequests)

	handleRequest(endpoint, method, action, serverName)
	fmt.Println("Request handling completed safely")

	fmt.Printf("Server %s is still running\n", serverName)
	fmt.Println("Server Summary:")
	fmt.Println("Name:", serverName)
	fmt.Println("Max Requests:", maxRequests)
	fmt.Println("Last Request:", method, endpoint)
	fmt.Println("Action Performed:", action)
	fmt.Println("Status: Operational")
}
