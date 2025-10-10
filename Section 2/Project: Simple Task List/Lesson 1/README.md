In this challenge, you'll create the foundation for a task management system by setting up the basic data structure and user interface. This is the first step in building a complete task list application that will grow throughout the project.

You will receive two inputs:

A string representing the application name (e.g., "MyTasks")
A string representing the user's name (e.g., "Alice")
Your task is to:

Define a Task struct with two fields:
Name (string) - the task description
Completed (bool) - whether the task is finished
Create a slice of Task structs to hold your task list (initially empty)
Display a welcome message in the format: "Welcome to [app_name], [user_name]!"
Display a menu with the following options, each on a separate line:
"1. Add Task"
"2. View Tasks"
"3. Complete Task"
"4. Remove Task"
"5. Exit"
Print the current number of tasks in the format: "Current tasks: [count]"
This challenge establishes the core structure that will support all future functionality in your task management system. The Task struct will be the building block for storing individual tasks, while the slice will manage your collection of tasks. The menu provides a clear interface for users to understand what actions they can perform.