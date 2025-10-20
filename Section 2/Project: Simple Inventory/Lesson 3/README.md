Extend your inventory management system by implementing functionality to add new products safely. This challenge builds on the previous stock checking system and adds the ability to expand your inventory while preventing duplicate entries through proper error handling.

You will receive two inputs:

A string containing existing inventory data in the format "product1:price1:quantity1,product2:price2:quantity2,product3:price3:quantity3" (e.g., "Laptop:999.99:5,Mouse:25.50:15,Keyboard:75.00:8")
A string containing new products to add in the format "product1:price1:quantity1,product2:price2:quantity2,product3:price3:quantity3" (e.g., "Monitor:299.99:10,Mouse:19.99:20,Webcam:89.99:12")
Your task is to:

Use the same Product struct from previous lessons with Price (float64) and Quantity (int) fields
Parse the first input by splitting on commas to get individual existing product entries
For each existing product entry, split on colons to get product name, price, and quantity
Convert the price string to float64 and quantity string to int
Create and populate the inventory map with the existing product data
Create a function called addNewItem that takes the inventory map, product name, price, and quantity as parameters and returns an error:
If the product already exists in the inventory, return an error with the message "product already exists: [product_name]"
If the product doesn't exist, add it to the inventory and return nil
Parse the second input by splitting on commas to get the list of new products to add
For each new product entry, split on colons to get product name, price, and quantity
Convert the price string to float64 and quantity string to int
Display the addition process header: "Adding New Items:"
For each new product, call the addNewItem function and display the results:
If no error: "[product_name]: Successfully added - $[price] (Stock: [quantity])"
If error: "[product_name]: Failed to add - [error_message]"
Count and display addition statistics:
"Addition Summary:"
"Items processed: [total_number_of_items_processed]"
"Items added: [number_of_items_successfully_added]"
"Items rejected: [number_of_items_that_failed_to_add]"
Display the updated inventory in alphabetical order by product name:
"Updated Inventory:"
For each product: "- [product_name]: $[price] (Stock: [quantity])"
Calculate and display final inventory statistics:
"Final Inventory Statistics:"
"Total Products: [total_number_of_products_in_inventory]"
"Total Items in Stock: [sum_of_all_quantities]"
"Total Inventory Value: $[sum_of_price_times_quantity_for_all_products]"
List any rejected products:
If there are rejected products: "Rejected products: [comma_separated_list_of_rejected_products]"
If no rejected products: "All new products were added successfully"
Use the strings package to split the input strings, the strconv package to convert strings to numbers, the errors package to create error messages, the fmt package for formatted output, and the sort package to sort product names alphabetically. Format all prices to 2 decimal places and the total inventory value to 2 decimal places. This challenge demonstrates how to safely expand data collections while preventing duplicate entries through proper error handling and validation.