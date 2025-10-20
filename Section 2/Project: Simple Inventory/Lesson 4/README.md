Enhance your inventory management system by implementing stock quantity updates with comprehensive validation and error handling. This challenge builds on the previous functionality and adds the ability to modify existing product quantities while preventing invalid operations that could result in negative stock levels.

You will receive two inputs:

A string containing existing inventory data in the format "product1:price1:quantity1,product2:price2:quantity2,product3:price3:quantity3" (e.g., "Laptop:999.99:5,Mouse:25.50:15,Keyboard:75.00:8")
A string containing stock update requests in the format "product1:change1,product2:change2,product3:change3" where changes can be positive or negative integers (e.g., "Mouse:10,Tablet:-5,Keyboard:-3")
Your task is to:

Use the same Product struct from previous lessons with Price (float64) and Quantity (int) fields
Parse the first input by splitting on commas to get individual existing product entries
For each existing product entry, split on colons to get product name, price, and quantity
Convert the price string to float64 and quantity string to int
Create and populate the inventory map with the existing product data
Create a function called updateStock that takes the inventory map, product name, and quantity change as parameters and returns an error:
If the product doesn't exist in the inventory, return an error with the message "product not found: [product_name]"
If the quantity change would result in a negative stock (current quantity + change < 0), return an error with the message "insufficient stock: cannot reduce [product_name] by [absolute_change_value], only [current_quantity] available"
If the update is valid, modify the product's quantity in the inventory and return nil
Parse the second input by splitting on commas to get the list of stock update requests
For each update request, split on colons to get product name and quantity change
Convert the quantity change string to int
Display the update process header: "Processing Stock Updates:"
For each update request, call the updateStock function and display the results:
If no error and change is positive: "[product_name]: Added [change] units - New stock: [new_quantity]"
If no error and change is negative: "[product_name]: Removed [absolute_change_value] units - New stock: [new_quantity]"
If no error and change is zero: "[product_name]: No change - Current stock: [current_quantity]"
If error: "[product_name]: Update failed - [error_message]"
Count and display update statistics:
"Update Summary:"
"Updates processed: [total_number_of_updates_processed]"
"Updates successful: [number_of_successful_updates]"
"Updates failed: [number_of_failed_updates]"
Display the updated inventory in alphabetical order by product name:
"Updated Inventory:"
For each product: "- [product_name]: $[price] (Stock: [quantity])"
Calculate and display final inventory statistics:
"Final Inventory Statistics:"
"Total Products: [total_number_of_products_in_inventory]"
"Total Items in Stock: [sum_of_all_quantities]"
"Total Inventory Value: $[sum_of_price_times_quantity_for_all_products]"
List any failed updates:
If there are failed updates: "Failed updates: [comma_separated_list_of_failed_product_names]"
If no failed updates: "All stock updates were processed successfully"
Use the strings package to split the input strings, the strconv package to convert strings to numbers, the errors package to create error messages, the fmt package for formatted output, and the sort package to sort product names alphabetically. Format all prices to 2 decimal places and the total inventory value to 2 decimal places. This challenge demonstrates how to safely modify existing data while preventing invalid operations through comprehensive validation and error handling.