package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	var fictionBooks string
	var booksToRemove string
	fmt.Scanln(&fictionBooks)
	fmt.Scanln(&booksToRemove)

	fictionList := strings.Split(fictionBooks, ",")
	removeList := strings.Split(booksToRemove, ",")

	bookSet := map[string]struct{}{}

	for _, book := range fictionList {
		bookSet[book] = struct{}{}
	}

	successfullyDelete := 0
	notFound := 0
	for _, removeBook := range removeList {
		if _, exists := bookSet[removeBook]; exists {
			delete(bookSet, removeBook)
			fmt.Println("Removing:", removeBook)
			successfullyDelete += 1
		} else {
			fmt.Println("Not found:", removeBook)
			notFound += 1
		}
	}

	fmt.Println("Initial collection size:", len(fictionList))
	fmt.Println("Removal requests:", len(removeList))
	fmt.Println("Books successfully removed:", successfullyDelete)
	fmt.Println("Books not found:", notFound)
	fmt.Println("Final collection size:", len(bookSet))
	fmt.Println("Remaining books in fiction collection:")

	if len(bookSet) == 0 {
		fmt.Println("- Collection is empty")
	} else {

		keys := make([]string, 0, len(bookSet))
		for k := range bookSet {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, book := range keys {
			fmt.Println("-", book)
		}
	}
}
