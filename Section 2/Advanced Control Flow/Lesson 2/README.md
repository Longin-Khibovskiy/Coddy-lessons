In this challenge, you'll practice using labeled break statements to exit nested loops when searching through a 2D grid. You need to find a specific target number in the grid and stop searching immediately once you find it.

You will receive two inputs:

A string representing the dimensions of the grid in the format "rows,cols" (e.g., "3,4")
A string representing the target number to search for (e.g., "7")
You are provided with the following 2D grid of numbers:

grid := [][]int{
{1, 2, 3, 4},
{5, 6, 7, 8},
{9, 10, 11, 12}
}
Your task is to:

Parse the dimensions from the first input to determine how many rows and columns to search
Parse the target number from the second input
Use nested loops with a labeled break to search through the grid
When you find the target number, print "Found [target] at position ([row], [col])" and immediately exit both loops using the labeled break
If you complete the search without finding the target, print "Target [target] not found"
The position should be reported using 0-based indexing (first row is 0, first column is 0).