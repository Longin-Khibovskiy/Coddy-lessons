package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var dbOperation string
	var appContext string
	fmt.Scanln(&dbOperation)
	fmt.Scanln(&appContext)

	dbParts := strings.Split(dbOperation, ",")
	databaseName := dbParts[0]
	operation := dbParts[1]
	errorType := dbParts[2]

	contextParts := strings.Split(appContext, ",")
	serviceName := contextParts[0]
	retryCount := contextParts[1]

	ErrConnectionTimeout := errors.New("connection timeout")
	ErrConnectionRefused := errors.New("connection refused")
	ErrDatabaseNotFound := errors.New("database not found")

	detectedType := ""
	var originalError error
	switch errorType {
	case "connection_timeout":
		originalError = ErrConnectionTimeout
		detectedType = "Connection Timeout"
	case "connection_refused":
		originalError = ErrConnectionRefused
		detectedType = "Connection Refused"
	case "database_not_found":
		originalError = ErrDatabaseNotFound
		detectedType = "Database Not Found"
	default:
		originalError = errors.New("unknown database error")
		detectedType = "Unknown Error"
	}

	DBerr := fmt.Errorf("failed to %s database %s: %w", operation, databaseName, originalError)

	serviceErr := fmt.Errorf("%s database operation failed after %s retries: %w", serviceName, retryCount, DBerr)

	fmt.Println("Checking for connection timeout:", errors.Is(serviceErr, ErrConnectionTimeout))
	fmt.Println("Checking for connection refused:", errors.Is(serviceErr, ErrConnectionRefused))
	fmt.Println("Checking for database not found:", errors.Is(serviceErr, ErrDatabaseNotFound))

	fmt.Println("Error Chain:")
	fmt.Println("Original:", originalError)
	fmt.Println("Database level:", DBerr)
	fmt.Println("Service level:", serviceErr)

	fmt.Println("Error Analysis:")
	fmt.Println("Database:", databaseName)
	fmt.Println("Operation:", operation)
	fmt.Println("Service:", serviceName)
	fmt.Println("Retries:", retryCount)
	fmt.Println("Error type detected:", detectedType)
}
