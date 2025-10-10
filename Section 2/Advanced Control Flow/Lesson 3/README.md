In this challenge, you'll practice using labeled continue statements to control nested loop execution. You're building a data processing system that needs to skip certain outer loop iterations based on conditions found in inner loops.

You will receive two inputs:

A string representing the number of rows to process (e.g., "4")
A string representing a skip condition number (e.g., "3")
You are provided with the following 2D data array:

data := [][]int{
{1, 2, 3, 4, 5},
{6, 7, 8, 9, 10},
{11, 12, 13, 14, 15},
{16, 17, 18, 19, 20},
{21, 22, 23, 24, 25}
}
Your task is to:

Parse the number of rows from the first input to determine how many rows to process
Parse the skip condition number from the second input
Use nested loops with a labeled continue to process the data
For each row, check each column value in that row
If you find the skip condition number in any column of a row, use labeled continue to skip the rest of that row's processing and move to the next row
For rows that don't contain the skip condition number, print "Processing row [row_number]: [value1] [value2] [value3] [value4] [value5]"
For rows that do contain the skip condition number, print "Skipping row [row_number] due to condition"
The row numbers should be reported using 0-based indexing (first row is 0). Process only the number of rows specified in the first input, and only check the first 5 columns of each row.