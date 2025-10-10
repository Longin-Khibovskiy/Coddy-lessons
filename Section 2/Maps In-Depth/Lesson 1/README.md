Create a student grade management system using maps with structs as values. This challenge demonstrates how to organize complex student data using meaningful keys for efficient lookup and retrieval.

You will receive two inputs:

A string representing the number of students (e.g., "4")
A string containing student data in the format: "name1:id1:grade1,name2:id2:grade2,name3:id3:grade3" (e.g., "Alice:101:A,Bob:102:B,Charlie:103:C,Diana:104:A")
Your task is to:

Define a Student struct with the following fields:
ID (int) - the student's ID number
Grade (string) - the student's letter grade
Create a map where the keys are student names (strings) and the values are Student structs
Parse the student data from the second input:
Split the input by commas to get individual student entries
For each entry, split by colons to separate name, ID, and grade
Convert the ID string to an integer
Create a Student struct and add it to the map with the name as the key
Display all students in alphabetical order by name in the format: "[name]: ID [id], Grade [grade]"
Calculate and display grade statistics:
Count how many students have each grade (A, B, C, D, F)
Display the count for each grade that appears in the format: "Grade [grade]: [count] students"
Display grades in alphabetical order (A, B, C, D, F)
Find and display the student with the highest ID in the format: "Highest ID: [name] ([id])"
Display the total number of students: "Total students: [count]"
Use the strconv package to convert the ID string to an integer, and the sort package to sort the student names alphabetically. This challenge demonstrates how maps with struct values provide an efficient way to organize and access complex data collections using meaningful identifiers.