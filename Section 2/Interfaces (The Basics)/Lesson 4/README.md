In this challenge, you'll practice working with the empty interface by creating a data processor that can handle different types of values. You'll create a function that accepts an interface{} parameter and processes various data types.

You will receive two inputs:

A string representing the data type (e.g., "int", "string", "bool", or "slice")
A string representing the actual value (e.g., "42", "hello", "true", or "1,2,3")
Your task is to:

Create a function called processData that takes an interface{} parameter and prints information about the value
Based on the data type input, convert the value string to the appropriate Go type:
For "int": convert to integer
For "string": use as string
For "bool": convert to boolean
For "slice": convert comma-separated values to a slice of integers
Call the processData function with the converted value
Inside processData, print the value and its type in the format: "Value: [value], Type: [type]"
For the slice input, the comma-separated string like "1,2,3" should be converted to []int{1, 2, 3}. The function should work with any type since it accepts interface{}, demonstrating how the empty interface can hold values of any type. Use fmt.Printf with the %v and %T format verbs to display the value and type respectively.