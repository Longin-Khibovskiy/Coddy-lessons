package main

import "fmt"

// calculateCalories determines how many calories are burned during an activity
// Parameters:
// - activity: the type of exercise ("running", "swimming", "cycling", or any other activity)
// - duration: how long the activity was performed (in minutes)
// - intensity: a multiplier representing how intense the workout was (1.0 is normal)
// Returns:
// - the total calories burned
func calculateCalories(activity string, duration int, intensity float64) float64 {
	// TODO: Implement this function to calculate calories burned based on:
	// 1. Determine the base calorie rate per minute based on activity type:
	//    - "running" burns 10 calories per minute at intensity 1.0
	//    - "swimming" burns 8 calories per minute at intensity 1.0
	//    - "cycling" burns 7 calories per minute at intensity 1.0
	//    - Any other activity burns 5 calories per minute at intensity 1.0
	// 2. Calculate total calories as: base calories × duration × intensity
	// 3. Return the calculated calories
	sum := 0.0
	if activity == "running" {
		sum = float64(duration) * float64(intensity) * 10.0
	} else if activity == "swimming" {
		sum = float64(duration) * float64(intensity) * 8.0
	} else if activity == "cycling" {
		sum = float64(duration) * float64(intensity) * 7.0
	} else {
		sum = float64(duration) * float64(intensity) * 5.0
	}
	return sum // Replace with your implementation
}

func main() {
	// Test the function with different activities
	fmt.Println("Running for 30 minutes at intensity 1.2:", calculateCalories("running", 30, 1.2))
	fmt.Println("Swimming for 45 minutes at intensity 1.0:", calculateCalories("swimming", 45, 1.0))
	fmt.Println("Cycling for 60 minutes at intensity 1.5:", calculateCalories("cycling", 60, 1.5))
	fmt.Println("Yoga for 60 minutes at intensity 0.8:", calculateCalories("yoga", 60, 0.8))
}
