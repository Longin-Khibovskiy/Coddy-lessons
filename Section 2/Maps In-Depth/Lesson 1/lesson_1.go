package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	ID    int
	Grade string
}

func main() {
	// Read input
	var numStudentsStr int
	var studentData string
	fmt.Scanln(&numStudentsStr)
	fmt.Scanln(&studentData)
	highStudenId := 0
	highStudenName := ""
	gradeA, gradeB, gradeC, gradeD, gradeE, gradeF := 0, 0, 0, 0, 0, 0
	students := make(map[string]Student)
	// TODO: Parse the student data and populate the map
	studentsData := strings.Split(studentData, ",")
	for i := numStudentsStr - 1; i >= 0; i-- {
		stud := strings.Split(studentsData[i], ":")
		id, _ := strconv.Atoi(stud[1])
		if id > highStudenId {
			highStudenId = id
			highStudenName = stud[0]
		}
		students[stud[0]] = Student{id, stud[2]}
		switch stud[2] {
		case "A":
			gradeA++
		case "B":
			gradeB++
		case "C":
			gradeC++
		case "D":
			gradeD++
		case "E":
			gradeE++
		case "F":
			gradeF++
		}
	}
	names := make([]string, 0, len(students))
	for name := range students {
		names = append(names, name)
	}
	sort.Strings(names)

	// 🔹 Вывод студентов в алфавитном порядке
	for _, name := range names {
		stud := students[name]
		fmt.Printf("%s: ID %d, Grade %s\n", name, stud.ID, stud.Grade)
	}
	if gradeA != 0 {
		fmt.Printf("Grade A: %d students\n", gradeA)
	}
	if gradeB != 0 {
		fmt.Printf("Grade B: %d students\n", gradeB)
	}
	if gradeC != 0 {
		fmt.Printf("Grade C: %d students\n", gradeC)
	}
	if gradeD != 0 {
		fmt.Printf("Grade D: %d students\n", gradeD)
	}
	if gradeE != 0 {
		fmt.Printf("Grade E: %d students\n", gradeE)
	}
	if gradeF != 0 {
		fmt.Printf("Grade F: %d students\n", gradeF)
	}
	fmt.Printf("Highest ID: %s (%d)\n", highStudenName, highStudenId)
	fmt.Printf("Total students: %d", len(students))
}
