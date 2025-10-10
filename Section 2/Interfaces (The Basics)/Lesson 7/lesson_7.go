package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

type Parrot struct {
	Name string
}

func (p Person) Speak() string {
	return fmt.Sprintf("Hello, I'm %s", p.Name)
}
func (p Parrot) Speak() string {
	return fmt.Sprintf("Squawk! %s says hello!", p.Name)
}
func makeAllSpeak(speakers []Speaker) {
	for _, s := range speakers {
		fmt.Println(s.Speak())
	}
}
func main() {
	// Read input
	var numSpeakersStr string
	var speakerTypesStr string
	var speakerNamesStr string

	fmt.Scanln(&numSpeakersStr)
	fmt.Scanln(&speakerTypesStr)
	fmt.Scanln(&speakerNamesStr)

	// Parse input
	numSpeakers, _ := strconv.Atoi(numSpeakersStr)
	speakerTypes := strings.Split(speakerTypesStr, ",")
	speakerNames := strings.Split(speakerNamesStr, ",")

	// TODO: Write your code below
	// 1. Define the Speaker interface
	// 2. Define Person and Parrot structs
	// 3. Implement Speak() methods for both structs
	// 4. Create makeAllSpeak function
	// 5. Create speakers based on input and store in a slice
	var speakers []Speaker
	for i := 0; i < numSpeakers; i++ {
		switch speakerTypes[i] {
		case "person":
			speakers = append(speakers, Person{Name: speakerNames[i]})
		case "parrot":
			speakers = append(speakers, Parrot{Name: speakerNames[i]})
		}
	}
	// 6. Call makeAllSpeak with your slice
	makeAllSpeak(speakers)
}
