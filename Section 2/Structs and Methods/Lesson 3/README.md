In this challenge, you'll practice creating methods with pointer receivers by building a simple bank account system. You'll create an Account struct and define methods that can modify the account balance using pointer receivers.

You will receive three inputs:

A string representing the account holder's name (e.g., "John Smith")
A string representing the initial account balance (e.g., "1000.50")
A string representing a transaction amount (e.g., "250.75")
Your task is to:

Define an Account struct with two fields: Name (string) and Balance (float64)
Create a method called deposit with a pointer receiver that adds the given amount to the account balance
Create a method called withdraw with a pointer receiver that subtracts the given amount from the account balance
Create a method called displayBalance with a value receiver that prints the account information in the format: "Account: [name], Balance: $[balance]"
Parse the inputs to create an Account instance
Call displayBalance to show the initial balance
Call deposit with the transaction amount
Call displayBalance to show the balance after deposit
Call withdraw with the transaction amount
Call displayBalance to show the final balance
The balance should be displayed with two decimal places precision. This challenge demonstrates how pointer receivers allow methods to modify the original struct data, unlike value receivers which work with copies.