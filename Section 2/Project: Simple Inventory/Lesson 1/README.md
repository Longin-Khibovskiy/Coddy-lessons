Set up the foundation for your inventory management system by creating the core data structures and displaying initial inventory data. This challenge establishes the Product struct and inventory map that will be used throughout the project.

You will receive two inputs:

A string containing store information in the format "store_name,location" (e.g., "TechStore,Downtown")
A string containing initial product data in the format "product1:price1:quantity1,product2:price2:quantity2,product3:price3:quantity3" (e.g., "Laptop:999.99:5,Mouse:25.50:15,Keyboard:75.00:8")
Your task is to:

Create a Product struct with two fields:
Price (float64) - the price of the product
Quantity (int) - the current stock quantity
Parse the first input by splitting on commas to get store name and location
Parse the second input by splitting on commas to get individual product entries
For each product entry, split on colons to get product name, price, and quantity
Convert the price string to float64 and quantity string to int
Create a map called inventory where keys are product names (strings) and values are Product structs
Populate the inventory map with the parsed product data
Display the store information:
"=== [store_name] Inventory System ==="
"Location: [location]"
"Inventory initialized with [number_of_products] products"
Display the current inventory in alphabetical order by product name:
"Current Inventory:"
For each product: "- [product_name]: $[price] (Stock: [quantity])"
Calculate and display inventory statistics:
"Inventory Statistics:"
"Total Products: [total_number_of_products]"
"Total Items in Stock: [sum_of_all_quantities]"
"Total Inventory Value: $[sum_of_price_times_quantity_for_all_products]"
Display system status:
"System Status: Ready"
"Inventory management system initialized successfully"
Use the strings package to split the input strings, the strconv package to convert strings to numbers, the fmt package for formatted output, and the sort package to sort product names alphabetically. Format the total inventory value to 2 decimal places. This challenge establishes the foundation data structures that will be extended in subsequent lessons of the inventory management project.