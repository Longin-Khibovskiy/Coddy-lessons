package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type APIError struct {
	Code    int
	Field   string
	Message string
}

func (apiError APIError) Error() string {
	return fmt.Sprintf("API error [%d]: %s (field: %s)", apiError.Code, apiError.Message, apiError.Field)
}

func main() {
	var apiDetails string
	var systemContext string
	fmt.Scanln(&apiDetails)
	fmt.Scanln(&systemContext)

	apiParts := strings.Split(apiDetails, ",")
	endpoint := apiParts[0]
	method := apiParts[1]
	errorCodeStr := apiParts[2]
	errorField := apiParts[3]

	contextParts := strings.Split(systemContext, ",")
	serviceName := contextParts[0]
	requestID := contextParts[1]

	errorCode, _ := strconv.Atoi(errorCodeStr)

	var apiError error
	switch errorCode {
	case 400:
		apiError = APIError{Code: 400, Field: errorField, Message: "invalid request data"}
	case 401:
		apiError = APIError{Code: 401, Field: errorField, Message: "authentication required"}
	case 404:
		apiError = APIError{Code: 404, Field: errorField, Message: "resource not found"}
	case 500:
		apiError = APIError{Code: 500, Field: errorField, Message: "internal server error"}
	default:
		apiError = APIError{Code: errorCode, Field: errorField, Message: "unknown error"}
	}
	endPointLevel := fmt.Errorf("%s %s failed: %w", method, endpoint, apiError)
	serviceLevel := fmt.Errorf("%s request %s failed: %w", serviceName, requestID, endPointLevel)
	var extractError APIError
	if errors.As(serviceLevel, &extractError) {
		fmt.Println("API error extracted successfully")
		fmt.Println("Extracted Error Details:")
		fmt.Println("Code:", extractError.Code)
		fmt.Println("Field:", extractError.Field)
		fmt.Println("Message:", extractError.Message)
		fmt.Println("Full error:", extractError.Error())
	} else {
		fmt.Println("Failed to extract API error")
	}

	fmt.Println("Error Chain:")
	fmt.Println("Original:", apiError.Error())
	fmt.Println("Endpoint level:", endPointLevel)
	fmt.Println("Service level:", serviceLevel)

	fmt.Println("Error Analysis:")
	fmt.Println("Service:", serviceName)
	fmt.Println("Request ID:", requestID)
	fmt.Println("Endpoint:", method, endpoint)
	fmt.Println("HTTP Status:", extractError.Code)
	fmt.Println("Problem Field:", extractError.Field)
	if extractError.Code >= 400 && extractError.Code <= 499 {
		fmt.Println("Error Category: Client Error")
	} else if extractError.Code >= 500 && extractError.Code <= 599 {
		fmt.Println("Error Category: Server Error")
	} else {
		fmt.Println("Error Category: Unknown")
	}
}
