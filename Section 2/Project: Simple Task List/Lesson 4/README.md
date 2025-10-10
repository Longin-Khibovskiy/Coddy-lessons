In this part of your task management system, you'll implement the ability to mark tasks as complete. This challenge focuses on using pointer receivers to modify the original task data, demonstrating how pointers allow you to update struct fields permanently.

Important Note: Since task names may contain spaces (like "Buy groceries"), you'll need to use bufio.Scanner to read entire lines of input. The fmt.Scanf function stops reading at the first space, which would break task names with spaces. Use bufio.NewScanner(os.Stdin) and call scanner.Scan() followed by scanner.Text() to read complete lines.

You will receive three inputs:

A string representing the number of tasks (e.g., "4")
A string containing task data in the format: "name1:status1,name2:status2,name3:status3" where status is either "true" or "false" (e.g., "Buy groceries:false,Call dentist:true,Finish project:false,Review code:false")
A string representing the index of the task to complete (e.g., "2" for the third task, using 0-based indexing)
Your task is to:

Define the same Task struct from previous challenges with Name (string) and Completed (bool) fields
Create a function called completeTask that takes a pointer to a slice of Task structs and an index (int) as parameters
Inside the completeTask function, modify the Completed field of the task at the given index to true
Use bufio.Scanner to read the input lines properly:
Create a scanner: scanner := bufio.NewScanner(os.Stdin)
For each input line: call scanner.Scan() then scanner.Text()
This ensures task names with spaces are read correctly
Parse the task data from the second input to create your slice of tasks:
Split the input by commas to get individual task entries
For each entry, split by colon to separate the name and status
Convert the status string to a boolean ("true" becomes true, "false" becomes false)
Convert the third input string to an integer to get the task index
Call your completeTask function with a pointer to your slice and the index
After completing the task, display all tasks using the same format as the previous challenge:
For completed tasks: "[x] [task_name]"
For incomplete tasks: "[ ] [task_name]"
Print a confirmation message: "Task '[task_name]' marked as completed!"
Print the updated summary: "Total: [total_count] tasks ([completed_count] completed, [incomplete_count] remaining)"
Use the strconv package to convert the index string to an integer. This challenge demonstrates how pointer receivers enable you to modify the original data structure, making permanent changes to your task list that persist after the function returns.