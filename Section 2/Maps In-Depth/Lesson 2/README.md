Build a simple inventory management system that demonstrates the power of using pointers to structs as map values. This challenge shows how pointer-based maps enable direct modification of stored data, making your inventory system efficient and responsive to updates.

You will receive three inputs:

A string representing the number of products (e.g., "4")
A string containing product data in the format: "name1:price1:quantity1,name2:price2:quantity2,name3:price3:quantity3" (e.g., "laptop:999.99:5,mouse:25.50:20,keyboard:75.00:15,monitor:299.99:8")
A string containing update operations in the format: "operation1:product1:value1,operation2:product2:value2" where operations are either "price" or "quantity" (e.g., "price:laptop:899.99,quantity:mouse:25,price:keyboard:65.00")
Your task is to:

Define a Product struct with the following fields:
Price (float64) - the product's price
Quantity (int) - the product's stock quantity
Create a map where the keys are product names (strings) and the values are pointers to Product structs
Parse the product data from the second input:
Split the input by commas to get individual product entries
For each entry, split by colons to separate name, price, and quantity
Convert the price string to a float64 and quantity string to an integer
Create a Product struct and store a pointer to it in the map
Display the initial inventory with the header "Initial Inventory:"
Show each product in the format: "[name]: $[price] (Stock: [quantity])"
Parse and apply the update operations from the third input:
Split the input by commas to get individual operations
For each operation, split by colons to get operation type, product name, and new value
Update the appropriate field directly through the map pointer
Print each update in the format: "Updated [product]: [field] changed to [new_value]"
Display the updated inventory with the header "Updated Inventory:"
Show each product again in the same format as the initial display
Calculate and display the total inventory value: "Total Inventory Value: $[total_value]"
Use the strconv package to convert price and quantity strings to their respective numeric types. Format all prices to 2 decimal places when displaying. This challenge demonstrates how storing pointers in maps allows you to modify the original struct data directly, making your inventory updates immediate and persistent without needing to reassign values to the map.