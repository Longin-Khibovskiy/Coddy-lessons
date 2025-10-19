Implement a stock checking function for your inventory management system that safely retrieves product quantities and handles cases where products don't exist. This challenge builds on the inventory foundation from the previous lesson and adds error handling for stock queries.

You will receive two inputs:

A string containing existing inventory data in the format "product1:price1:quantity1,product2:price2:quantity2,product3:price3:quantity3" (e.g., "Laptop:999.99:5,Mouse:25.50:15,Keyboard:75.00:8")
A string containing stock check requests in the format "product1,product2,product3" (e.g., "Mouse,Tablet,Keyboard")
Your task is to:

Use the same Product struct from the previous lesson with Price (float64) and Quantity (int) fields
Parse the first input by splitting on commas to get individual product entries
For each product entry, split on colons to get product name, price, and quantity
Convert the price string to float64 and quantity string to int
Create and populate the inventory map with the parsed product data
Create a function called checkStock that takes the inventory map and a product name as parameters and returns (int, error):
If the product exists in the inventory, return its quantity and nil
If the product doesn't exist, return 0 and an error with the message "product not found: [product_name]"
Parse the second input by splitting on commas to get the list of products to check
Display the stock checking header: "Stock Check Results:"
For each product in the check list, call the checkStock function and display the results:
If no error: "[product_name]: [quantity] units in stock"
If error: "[product_name]: Error - [error_message]"
Count and display summary statistics:
"Check Summary:"
"Products checked: [total_number_of_products_checked]"
"Products found: [number_of_products_that_exist]"
"Products not found: [number_of_products_that_dont_exist]"
Display the total stock for found products:
"Total stock for found products: [sum_of_quantities_for_existing_products] units"
List any missing products:
If there are missing products: "Missing products: [comma_separated_list_of_missing_products]"
If no missing products: "All requested products are available"
Use the strings package to split the input strings, the strconv package to convert strings to numbers, the errors package to create error messages, and the fmt package for formatted output. This challenge demonstrates how to implement safe data retrieval with proper error handling, a fundamental pattern in inventory management systems.