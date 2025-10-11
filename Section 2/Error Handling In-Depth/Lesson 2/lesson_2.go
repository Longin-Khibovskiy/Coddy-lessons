package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var fileDetails string
	var systemContext string
	fmt.Scanln(&fileDetails)
	fmt.Scanln(&systemContext)

	fileInfo := strings.Split(fileDetails, ",")
	filename := fileInfo[0]
	operation := fileInfo[1]
	status := fileInfo[2]

	contextInfo := strings.Split(systemContext, ",")
	serviceName := contextInfo[0]
	component := contextInfo[1]

	var originalErr error
	switch status {
	case "permission_denied":
		originalErr = errors.New("permission denied")
	case "not_found":
		originalErr = errors.New("file not found")
	case "corrupted":
		originalErr = errors.New("file corrupted")
	default:
		originalErr = errors.New("unknown error")
	}

	fileErr := fmt.Errorf("failed to %s file %s: %w", operation, filename, originalErr)
	componentErr := fmt.Errorf("%s error: %w", component, fileErr)
	serviceErr := fmt.Errorf("%s operation failed: %w", serviceName, componentErr)

	fmt.Println("Original error:", originalErr)
	fmt.Println("File level:", fileErr)
	fmt.Println("Component level:", componentErr)
	fmt.Println("Service level:", serviceErr)
	fmt.Println("Error Chain Summary:")
	fmt.Println("File:", filename)
	fmt.Println("Operation:", operation)
	fmt.Println("Component:", component)
	fmt.Println("Service:", serviceName)
	fmt.Println("Root cause:", originalErr)
}
