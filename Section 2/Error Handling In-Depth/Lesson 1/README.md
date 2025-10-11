Build a form validation system that uses custom error types to provide detailed feedback about validation failures. This challenge demonstrates how to create and use custom error types that implement the error interface to carry structured information about what went wrong.

You will receive two inputs:

A string containing user data in the format "username,email,age" (e.g., "john_doe,john@email.com,25")
A string containing validation rules in the format "min_username_length,max_age" (e.g., "5,65")
Your task is to:

Create a custom error type called ValidationError struct with two fields:
Field (string) - the name of the field that failed validation
Message (string) - a description of what went wrong
Implement the Error() string method for ValidationError that returns a formatted message: "Validation failed for [field]: [message]"
Create a validation function called validateUser that takes username, email, age, minimum username length, and maximum age as parameters
The validation function should check these rules and return a ValidationError if any fail:
Username must be at least the minimum length specified
Email must contain the "@" symbol
Age must be between 18 and the maximum age specified (inclusive)
Parse the input data by splitting the first input on commas to get username, email, and age
Parse the validation rules by splitting the second input on commas to get minimum username length and maximum age
Convert the age string to an integer and the rule strings to integers
Call the validation function with the parsed data
Display the validation results:
If validation passes: "User validation successful"
If validation fails: print the error message returned by the Error() method
After the validation result, display a summary:
"Validation Summary:"
"Username: [username]"
"Email: [email]"
"Age: [age]"
"Status: [PASSED or FAILED]"
Use the strings package to split the input strings and the strconv package to convert string numbers to integers. This challenge demonstrates how custom error types provide structured, meaningful error information that makes debugging and user feedback much more effective than generic error messages.