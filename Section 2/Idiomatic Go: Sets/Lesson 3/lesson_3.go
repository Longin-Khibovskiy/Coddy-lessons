package main

import (
	"fmt"
	"strings"
)

func main() {
	// Read input
	var teamMembersInput string
	var peopleToCheckInput string
	fmt.Scanln(&teamMembersInput)
	fmt.Scanln(&peopleToCheckInput)

	// Parse input strings
	teamMembers := strings.Split(teamMembersInput, ",")
	peopleToCheck := strings.Split(peopleToCheckInput, ",")

	members := map[string]struct{}{}
	found := 0
	notFound := 0
	for _, member := range teamMembers {
		members[member] = struct{}{}
	}
	for _, people := range peopleToCheck {
		if _, exists := members[people]; exists {
			fmt.Printf("%s is on the team\n", people)
			found += 1
		} else {
			fmt.Printf("%s is not on the team\n", people)
			notFound += 1
		}
	}
	fmt.Printf("Team size: %d\n", len(teamMembers))
	fmt.Printf("People checked: %d\n", len(peopleToCheck))
	fmt.Printf("Team members found: %d\nNon-team members: %d", found, notFound)
}
