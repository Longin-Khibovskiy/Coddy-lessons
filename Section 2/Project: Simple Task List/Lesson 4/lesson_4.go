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

func addTask(tasks []Task, name string, compl bool) []Task {
	return append(tasks, Task{name, compl})
}
func viewAllTasks(tasks []Task) {
	for _, task := range tasks {
		if task.Completed == true {
			fmt.Printf("[x] %s\n", task.Name)
		} else if task.Completed == false {
			fmt.Printf("[ ] %s\n", task.Name)
		}
	}
}

func main() {
	var completeds int
	var notcomleted int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numExistingTasks := scanner.Text()
	nums, _ := strconv.Atoi(strings.TrimSpace(numExistingTasks))
	var tasks []Task
	scanner.Scan()
	newInputTasks := scanner.Text()
	scanner.Scan()
	newCompletedTask := scanner.Text()
	num, _ := strconv.Atoi(strings.TrimSpace(newCompletedTask))
	var newTasks []string
	if newInputTasks != "" {
		newTasks = strings.Split(newInputTasks, ",")
	}
	recompletedName := ""
	for ini := range newTasks {
		completed := false
		name := strings.Split(newTasks[ini], ":")
		if ini == num {
			completed = true
			recompletedName = name[0]
			completeds += 1
		}
		if name[1] == "true" {
			completed = true
			completeds += 1
		} else {
			notcomleted += 1
		}
		tasks = addTask(tasks, name[0], completed)
	}
	viewAllTasks(tasks)
	fmt.Printf("Task '%s' marked as completed!\n", recompletedName)
	fmt.Printf("Total: %d tasks (%d completed, %d remaining)", nums, completeds, nums-completeds)
}
