In this challenge, you'll practice implementing both value and pointer receivers by creating a simple rectangle management system. You'll create a Rectangle struct and define two different types of methods that demonstrate when to use each receiver type.

You will receive three inputs:

A string representing the rectangle's width (e.g., "5.0")
A string representing the rectangle's height (e.g., "3.0")
A string representing a scaling factor (e.g., "2.0")
Your task is to:

Define a Rectangle struct with two fields: Width (float64) and Height (float64)
Create a method called area with a value receiver that calculates and returns the area of the rectangle (width × height)
Create a method called scale with a pointer receiver that multiplies both width and height by the given scaling factor
Parse the inputs to create a Rectangle instance
Print the initial area using the format: "Initial area: [area]"
Call the scale method with the scaling factor
Print the scaled area using the format: "Scaled area: [area]"
The area values should be displayed as whole numbers when they are integers (e.g., 15 instead of 15.0), but with decimal places when necessary (e.g., 15.5). This challenge demonstrates the key principle: use value receivers for methods that only read data (like calculating area) and pointer receivers for methods that modify the original struct (like scaling dimensions).