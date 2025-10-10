In this challenge, you'll practice using interface types by creating a drawing application that can work with different geometric shapes. You'll create a function that accepts an interface type as a parameter, demonstrating how interfaces make your code flexible and reusable.

You will receive three inputs:

A string representing the shape type (either "circle" or "rectangle")
A string representing the first dimension (radius for circle, width for rectangle)
A string representing the second dimension (not used for circle, height for rectangle)
Your task is to:

Define a Shape interface with one method signature: Area() that returns a float64
Define a Circle struct with a Radius field (float64)
Define a Rectangle struct with Width and Height fields (both float64)
Implement the Area() method for both structs:
Circle area: 3.14159 × radius × radius
Rectangle area: width × height
Create a function called printShapeInfo that takes a Shape interface as a parameter and prints the area in the format: "Area: [area]"
Based on the input shape type, create the appropriate shape instance and call printShapeInfo with it
The area should be displayed with up to 5 decimal places, but trailing zeros should be removed (e.g., 15.70796 instead of 15.70796000). This challenge demonstrates how a single function can work with multiple types through their shared interface, making your code more flexible and maintainable.