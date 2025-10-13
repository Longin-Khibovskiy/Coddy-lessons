Build an API error handling system that uses errors.As to extract and examine custom error types from wrapped error chains. This challenge demonstrates how to use errors.As to access the fields and data of specific custom error types, even when they've been wrapped with additional context.

You will receive two inputs:

A string containing API request details in the format "endpoint,method,error_code,error_field" (e.g., "/users,POST,400,email")
A string containing system context in the format "service_name,request_id" (e.g., "UserAPI,req_12345")
Your task is to:

Create a custom error type called APIError struct with three fields:
Code (int) - the HTTP status code
Field (string) - the field that caused the error
Message (string) - a description of the error
Implement the Error() string method for APIError that returns: "API error [code]: [message] (field: [field])"
Parse the first input by splitting on commas to get endpoint, method, error code, and error field
Parse the second input by splitting on commas to get service name and request ID
Convert the error code string to an integer
Create an APIError instance based on the error code:
If code is 400: message "invalid request data"
If code is 401: message "authentication required"
If code is 404: message "resource not found"
If code is 500: message "internal server error"
For any other code: message "unknown error"
Create an endpoint-level wrapped error using fmt.Errorf with %w: "[method] [endpoint] failed: %w"
Create a service-level wrapped error: "[service_name] request [request_id] failed: %w"
Use errors.As to extract the APIError from the wrapped error chain and display the extraction result:
If extraction succeeds: "API error extracted successfully"
If extraction fails: "Failed to extract API error"
If the extraction succeeded, display the extracted error details:
"Extracted Error Details:"
"Code: [code]"
"Field: [field]"
"Message: [message]"
"Full error: [full_error_message_from_Error_method]"
Display the complete error chain:
"Error Chain:"
"Original: [original_api_error_message]"
"Endpoint level: [endpoint_level_error_message]"
"Service level: [service_level_error_message]"
Display a final analysis:
"Error Analysis:"
"Service: [service_name]"
"Request ID: [request_id]"
"Endpoint: [method] [endpoint]"
"HTTP Status: [code]"
"Problem Field: [field]"
"Error Category: [error_category]" where error_category is "Client Error" for codes 400-499, "Server Error" for codes 500-599, or "Unknown" for other codes
Use the strings package to split the input strings on commas, the strconv package to convert the error code string to an integer, the errors package for errors.As, and the fmt package for error wrapping. This challenge demonstrates how errors.As allows you to access the specific fields and methods of custom error types within wrapped error chains, enabling sophisticated error handling and analysis.