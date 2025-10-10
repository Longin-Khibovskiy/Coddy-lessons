Build a library book management system that removes books from different genre collections using the Go set idiom. This challenge demonstrates how to use the delete() function to remove elements from sets while tracking the removal process.

You will receive two inputs:

A string containing books in the fiction collection separated by commas (e.g., "1984,Dune,Foundation,Neuromancer,Brave New World")
A string containing books to remove separated by commas (e.g., "Dune,Harry Potter,Foundation,Twilight,1984")
Your task is to:

Create a set using the Go idiom map[string]struct{} to store the fiction books
Parse the first input string by splitting it on commas to get individual book titles
Add each book to your set using the empty struct literal {} as the value
Parse the second input string by splitting it on commas to get the list of books to remove
For each book to remove, check if it exists in the set using the comma ok idiom
Display the removal process for each book being processed:
If the book exists in the collection: "Removing: [book_title]"
If the book doesn't exist in the collection: "Not found: [book_title]"
Use the delete() function to remove books that exist in the set
After processing all removal requests, display a summary:
"Initial collection size: [initial_count]"
"Removal requests: [requests_count]"
"Books successfully removed: [removed_count]"
"Books not found: [not_found_count]"
"Final collection size: [final_count]"
Finally, list all remaining books in the collection:
Header: "Remaining books in fiction collection:"
Each book on a separate line: "- [book_title]"
If no books remain: "- Collection is empty"
Use the strings package to split the input strings on commas. This challenge demonstrates how the delete() function safely removes elements from a Go set, with no errors thrown even when attempting to remove non-existent elements, making it perfect for cleanup operations.