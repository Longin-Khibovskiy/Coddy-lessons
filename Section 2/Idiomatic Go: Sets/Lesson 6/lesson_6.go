package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Read input
	var existingUsernames string
	var newRegistrations string
	fmt.Scanln(&existingUsernames)
	fmt.Scanln(&newRegistrations)

	// Parse input strings
	existing := strings.Split(existingUsernames, ",")
	attempts := strings.Split(newRegistrations, ",")

	// Create set using Go idiom
	userSet := make(map[string]struct{})

	// TODO: Write your code below
	// 1. Add existing usernames to the set
	for _, user := range existing {
		userSet[user] = struct{}{}
	}
	// 2. Process each registration attempt
	successful := 0
	notSuccessful := 0
	for _, user := range attempts {
		if _, exists := userSet[user]; exists {
			fmt.Println("Username taken:", user)
			notSuccessful += 1
		} else {
			userSet[user] = struct{}{}
			fmt.Println("Registered:", user)
			successful += 1
		}
	}
	fmt.Println("Registration Summary:")
	fmt.Println("Initial usernames:", len(existing))
	fmt.Println("Registration attempts:", len(attempts))
	fmt.Println("Successful registrations:", successful)
	fmt.Println("Rejected (duplicate):", notSuccessful)
	fmt.Println("Total registered usernames:", len(userSet))
	fmt.Println("All registered usernames:")
	keys := make([]string, 0, len(userSet))
	for k := range userSet {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("@%s\n", k)
	}
	// 3. Display registration results for each attempt
	// 4. Calculate and display the summary
	// 5. List all registered usernames (remember to sort for consistent output)

}
