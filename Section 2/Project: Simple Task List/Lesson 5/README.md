In this part of your task management system, you'll implement the ability to remove tasks from your list. This challenge focuses on slice manipulation using the append function to reconstruct a slice without a specific element, demonstrating how to permanently remove items from your data structure.

You will receive three inputs:

A string representing the number of tasks (e.g., "4")
A string containing task data in the format: "name1:status1,name2:status2,name3:status3" where status is either "true" or "false" (e.g., "Buy groceries:false,Call dentist:true,Finish project:false,Review code:true")
A string representing the index of the task to remove (e.g., "1" for the second task, using 0-based indexing)
Your task is to:

Define the same Task struct from previous challenges with Name (string) and Completed (bool) fields
Create a function called removeTask that takes a slice of Task structs and an index (int) as parameters
Inside the removeTask function, return a new slice with the task at the given index removed by using append to combine the elements before and after the target index
Parse the task data from the second input to create your slice of tasks:
Split the input by commas to get individual task entries
For each entry, split by colon to separate the name and status
Convert the status string to a boolean ("true" becomes true, "false" becomes false)
Convert the third input string to an integer to get the task index
Store the name of the task to be removed before calling the removeTask function
Call your removeTask function with your slice and the index, storing the returned slice
After removing the task, display all remaining tasks using the same format as previous challenges:
For completed tasks: "[x] [task_name]"
For incomplete tasks: "[ ] [task_name]"
Print a confirmation message: "Task '[task_name]' removed successfully!"
Print the updated summary: "Total: [total_count] tasks ([completed_count] completed, [incomplete_count] remaining)"
Use the strconv package to convert the index string to an integer. This challenge demonstrates how to use append to reconstruct slices by combining portions before and after a target element, effectively removing items from your collection while maintaining the order of remaining elements.