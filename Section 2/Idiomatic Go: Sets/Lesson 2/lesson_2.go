package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read input
	var countStr string
	var languagesStr string
	fmt.Scanln(&countStr)
	fmt.Scanln(&languagesStr)

	// Convert count string to integer (not needed for this challenge)
	count, _ := strconv.Atoi(countStr)
	if count == 0 {
		count = 1
	}

	// Split the languages string by commas
	languages := strings.Split(languagesStr, ",")

	// Create a set using map[string]struct{} idiom
	languageSet := map[string]struct{}{
		"Programming":     struct{}{},
		"Problem Solving": struct{}{},
		"Communication":   struct{}{},
	}

	// TODO: Write your code below
	// Process each language and check if it exists in the set
	// Use comma ok idiom: _, exists := languageSet[language]
	// Print "Added: [language]" or "Already exists: [language]"
	// Add languages to the set using empty struct literal {}
	countNewSkill := 0
	for _, language := range languages {
		if _, exists := languageSet[language]; exists {
			fmt.Printf("Already mastered: %s\n", language)
		} else {
			languageSet[language] = struct{}{}
			fmt.Printf("Learning new skill: %s\n", language)
			countNewSkill += 1
		}
	}

	fmt.Printf("Skills processed: %d\n", count)
	fmt.Printf("New skills learned: %d\n", countNewSkill)
	fmt.Printf("Total skills mastered: %d\n", len(languageSet))
	fmt.Println("Complete skill set:")

	// Convert map to slice and sort for consistent output
	keys := make([]string, 0, len(languageSet))
	for k := range languageSet {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// List each language
	for _, k := range keys {
		fmt.Println("✓", k)
	}
}
