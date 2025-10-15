Build a comprehensive mathematical calculator that implements safe division with multiple custom error types and proper error handling. This challenge combines custom error types, error wrapping, and error checking to create a robust system that handles various mathematical operation failures gracefully.

You will receive two inputs:

A string containing operation details in the format "operation,num1,num2,precision" (e.g., "divide,10,0,2")
A string containing calculator settings in the format "calculator_name,mode" (e.g., "ScientificCalc,strict")
Your task is to:

Create two custom error types:
DivisionError struct with fields Dividend (float64) and Divisor (float64)
ValidationError struct with fields Field (string) and Value (float64)
Implement the Error() string method for DivisionError that returns: "division error: cannot divide [dividend] by [divisor]"
Implement the Error() string method for ValidationError that returns: "validation error: invalid [field] value [value]"
Parse the first input by splitting on commas to get operation, num1, num2, and precision
Parse the second input by splitting on commas to get calculator name and mode
Convert the string numbers to float64 and precision to int
Create a safeDivide function that takes two float64 parameters and returns (float64, error):
If divisor is 0, return 0 and a DivisionError with the dividend and divisor values
If dividend is negative and mode is "strict", return 0 and a ValidationError with field "dividend" and the dividend value
If divisor is negative and mode is "strict", return 0 and a ValidationError with field "divisor" and the divisor value
Otherwise, return the division result and nil
Create a performCalculation function that wraps the safeDivide result:
If safeDivide returns an error, wrap it with: "calculation failed in [calculator_name]: %w"
If successful, return the result and nil
Call performCalculation and handle the result:
If no error: print "Calculation successful: [result_formatted_to_precision_decimal_places]"
If error exists: print "Calculation failed: [error_message]"
Use errors.As to check for specific error types in the wrapped error:
Print "Checking for division error: [true/false]"
Print "Checking for validation error: [true/false]"
If a DivisionError is found, print its details:
"Division Error Details:"
"Dividend: [dividend]"
"Divisor: [divisor]"
If a ValidationError is found, print its details:
"Validation Error Details:"
"Field: [field]"
"Value: [value]"
Display a final summary:
"Calculator Summary:"
"Name: [calculator_name]"
"Mode: [mode]"
"Operation: [operation]"
"Input: [num1] [operation_symbol] [num2]" where operation_symbol is "/" for divide
"Precision: [precision] decimal places"
"Status: [Success/Failed]"
Use the strings package to split the input strings on commas, the strconv package to convert strings to numbers, the errors package for errors.As, and the fmt package for error wrapping and formatting. This challenge demonstrates how custom error types, error wrapping, and proper error checking work together to create robust error handling systems that provide detailed information about different types of failures.