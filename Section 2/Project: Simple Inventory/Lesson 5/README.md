Complete your inventory management system by implementing a comprehensive reporting function that provides detailed insights into your current inventory. This challenge builds on all previous functionality and adds the ability to generate formatted reports with statistics, categorization, and analysis.

You will receive two inputs:

A string containing existing inventory data in the format "product1:price1:quantity1,product2:price2:quantity2,product3:price3:quantity3" (e.g., "Laptop:999.99:5,Mouse:25.50:15,Keyboard:75.00:8")
A string containing report configuration in the format "report_type,threshold_value" where report_type can be "full", "low_stock", or "high_value" and threshold_value is a number (e.g., "low_stock,10")
Your task is to:

Use the same Product struct from previous lessons with Price (float64) and Quantity (int) fields
Parse the first input by splitting on commas to get individual existing product entries
For each existing product entry, split on colons to get product name, price, and quantity
Convert the price string to float64 and quantity string to int
Create and populate the inventory map with the existing product data
Create a function called generateReport that takes the inventory map, report type, and threshold value as parameters
Parse the second input by splitting on commas to get report type and threshold value
Convert the threshold value string to float64
Display the report header based on the report type:
For "full": "=== FULL INVENTORY REPORT ==="
For "low_stock": "=== LOW STOCK REPORT ==="
For "high_value": "=== HIGH VALUE REPORT ==="
Display general inventory statistics:
"Total Products: [total_number_of_products]"
"Total Items in Stock: [sum_of_all_quantities]"
"Total Inventory Value: $[sum_of_price_times_quantity_for_all_products]"
"Average Product Price: $[total_inventory_value_divided_by_total_items]"
Filter and display products based on report type:
For "full": Show all products
For "low_stock": Show only products with quantity less than or equal to threshold value
For "high_value": Show only products with total value (price × quantity) greater than or equal to threshold value
Display the filtered products section:
For "full": "All Products:"
For "low_stock": "Products with stock ≤ [threshold_value]:"
For "high_value": "Products with value ≥ $[threshold_value]:"
For each filtered product (in alphabetical order by product name):
"- [product_name]: $[price] × [quantity] = $[total_value]"
Display filtered statistics:
"Filtered Results:"
"Products shown: [number_of_products_in_filtered_results]"
"Items in filtered products: [sum_of_quantities_for_filtered_products]"
"Value of filtered products: $[sum_of_values_for_filtered_products]"
Find and display the most expensive and least expensive products:
"Price Analysis:"
"Most expensive: [product_name] at $[price]"
"Least expensive: [product_name] at $[price]"
Find and display the highest and lowest stock products:
"Stock Analysis:"
"Highest stock: [product_name] with [quantity] units"
"Lowest stock: [product_name] with [quantity] units"
Display inventory health assessment:
Count products with quantity ≤ 5: "Low stock items (≤5): [count]"
Count products with quantity > 20: "High stock items (>20): [count]"
Count products with value ≥ $500: "High value items (≥$500): [count]"
Display report completion:
"Report generated successfully"
"Threshold applied: [threshold_value]"
Use the strings package to split the input strings, the strconv package to convert strings to numbers, the fmt package for formatted output, and the sort package to sort product names alphabetically. Format all prices and values to 2 decimal places. When finding most/least expensive or highest/lowest stock, if there are ties, use the alphabetically first product name. This challenge demonstrates how to create comprehensive reporting systems that provide valuable business insights through data analysis and formatted presentation.