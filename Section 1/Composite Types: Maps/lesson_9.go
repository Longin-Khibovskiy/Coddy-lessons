package main

import (
	"fmt"
	"sort"
)

func main() {
	// A map of student grades
	studentGrades := map[string]string{
		"Alice":   "A",
		"Bob":     "B",
		"Charlie": "B+",
		"David":   "A-",
		"Eva":     "C",
	}

	// Collect the student names (keys) into a slice to sort them
	var names []string
	for name := range studentGrades {
		names = append(names, name)
	}
	sort.Strings(names)

	// TODO: Iterate through the sorted list of student names
	// and print each student's name and grade in the format:
	// "Student: [name], Grade: [grade]"
	for _, student := range names {
		fmt.Printf("Student: %v, Grade: %s\n", student, studentGrades[student])
	}
}
