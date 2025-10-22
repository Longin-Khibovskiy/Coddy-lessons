# Неправильно сгенерированы выводы по заданию
### Сделано для вывода (не по заданию)

Create a custom mathutils package with basic mathematical operations and then use it from your main program. This challenge will test your ability to create a package, define exported functions, and import your custom package for use in another program.

You will receive two inputs:

A string containing function definitions in the format "function1:param1,param2:operation,function2:param1,param2:operation" (e.g., "Add:5,3:addition,Multiply:4,7:multiplication,Subtract:10,4:subtraction")
A string containing calculation requests in the format "operation1:value1,value2|operation2:value1,value2|operation3:value1,value2" (e.g., "Add:15,25|Multiply:6,8|Subtract:50,20")
Your task is to:

Create a mathutils package by declaring package mathutils at the top
Parse the first input by splitting on commas to get individual function definitions
For each function definition, split on colons to get function name, parameters, and operation type
Create the following exported functions in your mathutils package:
Add(a, b int) int - returns the sum of a and b
Subtract(a, b int) int - returns a minus b
Multiply(a, b int) int - returns the product of a and b
Divide(a, b int) int - returns a divided by b (integer division)
In your main program, import the mathutils package
Display the package information header: "=== MATHUTILS PACKAGE DEMO ==="
Parse the first input and display the available functions:
"Available functions in mathutils package:"
For each function: "- [function_name]: Performs [operation_type] on two integers"
Parse the second input by splitting on pipes to get individual calculation requests
For each calculation request, split on colons to get the function name and parameters
Split the parameters on commas to get the two integer values
Convert the parameter strings to integers using strconv.Atoi
Display the calculation process header: "Performing calculations:"
For each calculation request:
Call the appropriate function from the mathutils package
Display: "mathutils.[function_name]([param1], [param2]) = [result]"
Calculate and display summary statistics:
"Calculation Summary:"
"Total calculations performed: [number_of_calculations]"
"Functions used: [comma_separated_list_of_unique_function_names_used]"
"Sum of all results: [sum_of_all_calculation_results]"
Display package usage confirmation:
"Package demonstration completed successfully"
"All functions from mathutils package executed correctly"
Use the strings package to split the input strings, the strconv package to convert strings to integers, and the fmt package for formatted output. Remember that only functions starting with capital letters will be exported and accessible from other packages. This challenge demonstrates how to create reusable code through custom packages and how to organize functionality into logical modules.