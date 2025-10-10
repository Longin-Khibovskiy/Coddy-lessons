In this challenge, you'll practice implementing an interface by creating a simple vehicle system. You'll create a Car struct that automatically implements a Vehicle interface through Go's implicit implementation.

You will receive two inputs:

A string representing the car's brand (e.g., "Toyota")
A string representing the car's speed in km/h (e.g., "120")
Your task is to:

Define a Vehicle interface with one method signature: GetInfo() that returns a string
Define a Car struct with two fields: Brand (string) and Speed (int)
Implement the GetInfo() method for the Car struct that returns a string in the format: "Brand: [brand], Speed: [speed] km/h"
Parse the inputs to create a Car instance
Call the GetInfo() method and print the result
Remember that in Go, the Car struct will automatically implement the Vehicle interface because it has the required GetInfo() method with the correct signature. There's no need to explicitly declare that Car implements Vehicle - Go recognizes this relationship implicitly.