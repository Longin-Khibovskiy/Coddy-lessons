Build a file processing system that wraps errors with contextual information as they propagate through different layers of your application. This challenge demonstrates how to use the %w format verb with fmt.Errorf to create error chains that preserve the original error while adding meaningful context.

You will receive two inputs:

A string containing file operation details in the format "filename,operation,status" (e.g., "config.json,read,permission_denied")
A string containing system context in the format "service_name,component" (e.g., "UserService,ConfigLoader")
Your task is to:

Parse the first input by splitting on commas to get the filename, operation, and status
Parse the second input by splitting on commas to get the service name and component
Create an original error based on the status using errors.New():
If status is "permission_denied": "permission denied"
If status is "not_found": "file not found"
If status is "corrupted": "file corrupted"
For any other status: "unknown error"
Create a file-level wrapped error using fmt.Errorf with %w: "failed to [operation] file [filename]: %w"
Create a component-level wrapped error: "[component] error: %w"
Create a service-level wrapped error: "[service_name] operation failed: %w"
Display the error chain by printing each level:
"Original error: [original_error_message]"
"File level: [file_level_error_message]"
"Component level: [component_level_error_message]"
"Service level: [service_level_error_message]"
Display a summary of the error chain:
"Error Chain Summary:"
"File: [filename]"
"Operation: [operation]"
"Component: [component]"
"Service: [service_name]"
"Root cause: [original_error_message]"
Use the strings package to split the input strings on commas, the errors package to create the original error, and the fmt package to create wrapped errors. This challenge demonstrates how error wrapping with %w creates a clear trail of context as errors move through different layers of your application, making debugging much more effective.