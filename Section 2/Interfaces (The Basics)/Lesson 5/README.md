In this challenge, you'll practice using type assertions to safely extract and work with concrete values from interface variables. You'll create a data analyzer that processes different types of values stored in empty interfaces.

You will receive two inputs:

A string representing the data type to check for (e.g., "int", "string", or "bool")
A string representing the actual value (e.g., "42", "hello", or "true")
Your task is to:

Convert the value string to the appropriate Go type and store it in an interface{} variable:
For "int": convert to integer
For "string": use as string
For "bool": convert to boolean
Use type assertion to check if the interface variable contains the expected type specified in the first input
If the type assertion succeeds, print: "Success: [value] is a [type]"
If the type assertion fails, print: "Failed: value is not a [type]"
You must use the safe type assertion syntax with the ok variable to check if the assertion was successful. The challenge tests your understanding of how to safely extract concrete values from interface variables without causing your program to panic.

For boolean conversion, "true" should convert to true and any other string should convert to false. Use the strconv package for string to integer conversion.