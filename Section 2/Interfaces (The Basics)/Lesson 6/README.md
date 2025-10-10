In this challenge, you'll practice using type switches to handle multiple data types in a notification system. You'll create a function that processes different types of notifications and formats them appropriately based on their concrete type.

You will receive two inputs:

A string representing the notification type (e.g., "email", "sms", "push", or "alert")
A string representing the notification content or value
Your task is to:

Create a function called processNotification that takes an interface{} parameter
Based on the notification type input, convert the content to the appropriate Go type and store it in an interface{} variable:
For "email": use the content as a string
For "sms": convert the content to an integer (representing character count)
For "push": convert the content to a boolean (true for "enabled", false for anything else)
For "alert": convert the content to a float64 (representing priority level)
Call the processNotification function with the converted value
Inside processNotification, use a type switch to handle each type and print the appropriate message:
For string: "Email notification: [value]"
For int: "SMS notification with [value] characters"
For bool: "Push notifications: [value]"
For float64: "Alert with priority: [value]"
For any other type: "Unknown notification type"
Use the strconv package for string conversions. For boolean conversion, only "enabled" should result in true. This challenge demonstrates how type switches provide a clean way to handle multiple types in a single control structure, automatically giving you the correctly typed value in each case.