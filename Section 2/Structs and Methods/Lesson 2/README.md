In this challenge, you'll practice working with value receivers by creating a simple temperature monitoring system. You'll create a Sensor struct and define methods that demonstrate how value receivers work with copies of the struct.

You will receive three inputs:

A string representing the sensor ID (e.g., "TEMP001")
A string representing the current temperature reading (e.g., "23.5")
A string representing a temperature adjustment value (e.g., "2.0")
Your task is to:

Define a Sensor struct with two fields: ID (string) and Temperature (float64)
Create a method called displayReading with a value receiver that prints the sensor information in the format: "Sensor [ID]: [temperature]°C"
Create a method called adjustTemperature with a value receiver that adds the adjustment value to the temperature and prints: "Adjusted reading: [new_temperature]°C"
Parse the inputs to create a Sensor instance
Call the displayReading method to show the original temperature
Call the adjustTemperature method with the adjustment value
Call the displayReading method again to demonstrate that the original struct remains unchanged
This challenge will demonstrate that methods with value receivers work on copies of the struct, so the original struct's temperature value will remain unchanged even after calling the adjustTemperature method. The temperature values should be displayed with one decimal place precision.