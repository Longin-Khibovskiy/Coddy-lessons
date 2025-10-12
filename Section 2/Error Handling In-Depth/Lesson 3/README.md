Build a database connection monitoring system that uses errors.Is to identify specific connection errors in wrapped error chains. This challenge demonstrates how to use errors.Is to detect sentinel errors even when they've been wrapped with additional context through multiple application layers.

You will receive two inputs:

A string containing database operation details in the format "database_name,operation,error_type" (e.g., "users_db,connect,connection_timeout")
A string containing application context in the format "service_name,retry_count" (e.g., "AuthService,3")
Your task is to:

Define three sentinel errors using errors.New():
ErrConnectionTimeout with message "connection timeout"
ErrConnectionRefused with message "connection refused"
ErrDatabaseNotFound with message "database not found"
Parse the first input by splitting on commas to get database name, operation, and error type
Parse the second input by splitting on commas to get service name and retry count
Create the original error based on the error type:
If error type is "connection_timeout": use ErrConnectionTimeout
If error type is "connection_refused": use ErrConnectionRefused
If error type is "database_not_found": use ErrDatabaseNotFound
For any other error type: use errors.New("unknown database error")
Create a database-level wrapped error using fmt.Errorf with %w: "failed to [operation] database [database_name]: %w"
Create a service-level wrapped error: "[service_name] database operation failed after [retry_count] retries: %w"
Use errors.Is to check if the final wrapped error matches each sentinel error and display the results:
"Checking for connection timeout: [true/false]"
"Checking for connection refused: [true/false]"
"Checking for database not found: [true/false]"
Display the error chain:
"Error Chain:"
"Original: [original_error_message]"
"Database level: [database_level_error_message]"
"Service level: [service_level_error_message]"
Display the error analysis:
"Error Analysis:"
"Database: [database_name]"
"Operation: [operation]"
"Service: [service_name]"
"Retries: [retry_count]"
"Error type detected: [detected_error_type]" where detected_error_type is one of "Connection Timeout", "Connection Refused", "Database Not Found", or "Unknown Error"
Use the strings package to split the input strings on commas, the errors package for sentinel errors and errors.Is, and the fmt package for error wrapping. This challenge demonstrates how errors.Is provides robust error identification that works correctly with wrapped errors, making it essential for reliable error handling in complex applications.