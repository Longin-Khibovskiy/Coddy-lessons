package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// TODO: Define your Book struct here
type Book struct {
	title  string
	author string
	pages  int
}

// TODO: Define your displayInfo method here
func displayInfo(b Book) {
	fmt.Printf("Title: %s, Author: %s, Pages: %d\n", b.title, b.author, b.pages)
}

// TODO: Define your getDescription method here
func getDescription(b Book) {
	fmt.Printf("%s by %s", b.title, b.author)
}

func main() {
	// Create scanner for reading input lines
	scanner := bufio.NewScanner(os.Stdin)

	// Read title
	scanner.Scan()
	title := scanner.Text()

	// Read author
	scanner.Scan()
	author := scanner.Text()

	// Read pages
	scanner.Scan()
	pagesStr := scanner.Text()

	// Convert pages string to integer
	pages, _ := strconv.Atoi(pagesStr)

	// TODO: Create a Book instance and call the methods
	Book := Book{
		title:  title,
		author: author,
		pages:  pages,
	}
	displayInfo(Book)
	getDescription(Book)
}
