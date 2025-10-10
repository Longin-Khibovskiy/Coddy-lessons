Continuing with your task management system, you'll now implement the functionality to display all tasks in a user-friendly format. This challenge focuses on iterating through your task collection and presenting each task with its completion status using clear visual indicators.

You will receive two inputs:

A string representing the number of tasks (e.g., "4")
A string containing task data in the format: "name1:status1,name2:status2,name3:status3" where status is either "true" or "false" (e.g., "Buy groceries:false,Call dentist:true,Finish project:false,Review code:true")
Your task is to:

Define the same Task struct from previous challenges with Name (string) and Completed (bool) fields
Create a function called viewAllTasks that takes a slice of Task structs as a parameter
Inside the viewAllTasks function, iterate through the slice and print each task with the following format:
For completed tasks: "[x] [task_name]"
For incomplete tasks: "[ ] [task_name]"
Parse the task data from the second input to create your slice of tasks:
Split the input by commas to get individual task entries
For each entry, split by colon to separate the name and status
Convert the status string to a boolean ("true" becomes true, "false" becomes false)
Call your viewAllTasks function with the created slice
After displaying all tasks, print a summary in the format: "Total: [total_count] tasks ([completed_count] completed, [incomplete_count] remaining)"
For example, if you have tasks "Buy groceries:false,Call dentist:true,Finish project:false", your output should show each task with appropriate checkboxes and then display the summary with the total count and completion statistics. This challenge practices slice iteration, conditional formatting with if statements, and provides users with a clear view of their task progress.