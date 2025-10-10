Building on your task management foundation, you'll now implement the core functionality to add new tasks to your system. This challenge focuses on creating a function that expands your task list by appending new tasks to the existing slice.

Important Note: Since task names may contain spaces (like "Buy groceries"), you'll need to use bufio.Scanner to read entire lines of input. The fmt.Scanln function stops reading at the first space, which would break task names with spaces. Use bufio.NewScanner(os.Stdin) and call scanner.Scan() followed by scanner.Text() to read complete lines.

You will receive two inputs:

A string representing the number of existing tasks (e.g., "2")
A string containing task names separated by commas (e.g., "Buy groceries,Call dentist,Finish project")
Your task is to:

Define the same Task struct from the previous challenge with Name (string) and Completed (bool) fields
Create a function called addTask that takes a slice of Task structs and a task name (string) as parameters
The addTask function should return a new slice with the added task (the new task should have Completed set to false)
Use bufio.Scanner to read the input lines properly to handle task names with spaces
Create an initial slice of tasks based on the number of existing tasks specified in the first input:
If the number is "0", start with an empty slice
If the number is greater than "0", create that many placeholder tasks with names like "Existing Task 1", "Existing Task 2", etc., all marked as incomplete
Parse the comma-separated task names from the second input
Use your addTask function to add each new task to the slice
Print each task in the final slice in the format: "Task: [name], Completed: [true/false]"
Print the total number of tasks in the format: "Total tasks: [count]"
For example, if you have 1 existing task and want to add "Buy groceries,Call dentist", your output should show the existing task first, followed by the two new tasks, and finally the total count. This challenge practices slice manipulation with the append function and demonstrates how to build functionality that extends your data structures.