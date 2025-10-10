Create a safe map initialization system that prevents panic errors by checking for nil maps before attempting to write operations. This challenge demonstrates the critical importance of validating map state before performing operations that could crash your program.

You will receive two inputs:

A string representing the number of operations (e.g., "5")
A string containing operations in the format: "operation1:key1:value1,operation2:key2:value2" where operations are either "check", "init", or "add" (e.g., "check:scores,init:scores,add:Alice:95,add:Bob:87,check:grades")
Your task is to:

Create two map variables of type map[string]int named scores and grades without initializing them (they will be nil)
Process each operation from the input string sequentially:
For "check" operations:
Check if the specified map is nil
Print "Map [map_name] is nil" if the map is nil
Print "Map [map_name] is initialized" if the map is not nil
For "init" operations:
Initialize the specified map using make(map[string]int)
Print "Initialized map [map_name]"
For "add" operations:
First check if the map is nil
If the map is nil, print "Cannot add to nil map [map_name]"
If the map is not nil, add the key-value pair and print "Added [key]:[value] to [map_name]"
Parse the operations by splitting the input string:
Split by commas to get individual operations
For each operation, split by colons to get the operation type and parameters
Convert the value string to an integer for add operations
After processing all operations, display the final state of both maps:
For each map, if it's nil: "Final state - [map_name]: nil"
For each map, if it's initialized: "Final state - [map_name]: [count] entries" where count is the number of key-value pairs
Use the strconv package to convert value strings to integers for add operations. This challenge demonstrates how checking for nil maps before write operations prevents runtime panics and makes your code robust and reliable.