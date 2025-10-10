package main

import "fmt"

type Task struct {
	Name      string
	Completed bool
}

func main() {
	// Read input
	var appName string
	var userName string
	fmt.Scanln(&appName)
	fmt.Scanln(&userName)
	var tasks []Task
	fmt.Printf("Welcome to %s, %s!\n", appName, userName)
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Complete Task")
	fmt.Println("4. Remove Task")
	fmt.Println("5. Exit")
	fmt.Println("Current tasks:", len(tasks))
	// TODO: Write your code below
	// Define the Task struct
	// Create a slice of Task structs
	// Display welcome message and menu
	// Print current tasks count

}
