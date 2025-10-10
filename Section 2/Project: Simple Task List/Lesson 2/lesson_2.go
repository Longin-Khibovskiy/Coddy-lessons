package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Name      string
	Completed bool
}

func addTask(tasks []Task, name string) []Task {
	return append(tasks, Task{name, false})
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numExistingTasks := scanner.Text()
	num, _ := strconv.Atoi(strings.TrimSpace(numExistingTasks))
	var tasks []Task
	for i := 1; i <= num; i++ {
		tasks = append(tasks, Task{fmt.Sprintf("Existing Task %d", i), false})
	}

	scanner.Scan()
	newInputTasks := scanner.Text()
	var newTasks []string
	if newInputTasks != "" {
		newTasks = strings.Split(newInputTasks, ",")
	}
	for _, i := range newTasks {
		tasks = addTask(tasks, strings.TrimSpace(i))
	}

	for _, task := range tasks {
		fmt.Printf("Task: %s, Completed: %t\n", task.Name, task.Completed)
	}
	fmt.Println("Total tasks:", len(tasks))
}
