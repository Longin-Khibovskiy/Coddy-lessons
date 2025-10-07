package main

import "fmt"

func main() {
	// TODO: Create an anonymous struct variable called 'book' with fields:
	// - title (string): "The Go Programming Language"
	// - pages (int): 380
	book := struct {
		title string
		pages int
	}{
		title: "The Go Programming Language",
		pages: 380,
	}

	// Print the book information
	fmt.Printf("Book: %s, Pages: %d\n", book.title, book.pages)
}
