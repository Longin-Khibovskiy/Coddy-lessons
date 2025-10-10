In this challenge, you'll practice defining methods on structs by creating a simple book management system. You'll create a Book struct and define methods that allow the book to display its information.

Important Note: Since book titles and author names may contain spaces, you'll need to use bufio.Scanner to read entire lines of input. The fmt.Scanln function stops reading at the first space, which would break multi-word titles and author names. Use bufio.NewScanner(os.Stdin) and call scanner.Scan() followed by scanner.Text() to read complete lines.

You will receive three inputs:

A string representing the book title (e.g., "The Go Programming Language")
A string representing the author name (e.g., "Alan Donovan")
A string representing the number of pages (e.g., "380")
Your task is to:

Define a Book struct with three fields: Title (string), Author (string), and Pages (int)
Create a method called displayInfo with a value receiver that prints the book information in the format: "Title: [title], Author: [author], Pages: [pages]"
Create a method called getDescription with a value receiver that returns a string in the format: "[title] by [author]"
Use bufio.Scanner to read the input lines properly:
Create a scanner: scanner := bufio.NewScanner(os.Stdin)
For each input line: call scanner.Scan() then scanner.Text()
This ensures titles and author names with spaces are read correctly
Parse the inputs to create a Book instance
Call the displayInfo method to print the book information
Call the getDescription method and print the returned description
The methods should use the receiver variable to access the struct fields and perform their respective operations. Remember that methods with value receivers work with a copy of the struct instance.