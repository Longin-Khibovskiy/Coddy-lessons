package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var rockSongs string
	var popSongs string

	// Считываем через Scanln — как требует платформа
	fmt.Scanln(&rockSongs)
	fmt.Scanln(&popSongs)

	// Создаём множества
	rockSet := make(map[string]struct{})
	popSet := make(map[string]struct{})

	// Разделяем по запятой
	rockList := strings.Split(rockSongs, ",")
	popList := strings.Split(popSongs, ",")

	// Добавляем песни в множества
	for _, s := range rockList {
		s = strings.TrimSpace(s)
		if s != "" {
			rockSet[s] = struct{}{}
		}
	}
	for _, s := range popList {
		s = strings.TrimSpace(s)
		if s != "" {
			popSet[s] = struct{}{}
		}
	}

	// --- Rock ---
	fmt.Println("Rock Playlist Analysis:")
	fmt.Printf("Total rock songs: %d\n", len(rockSet))
	fmt.Println("Songs in rock playlist:")
	for _, s := range getSortedKeys(rockSet) {
		fmt.Printf("♪ %s\n", s)
	}

	// --- Pop ---
	fmt.Println("Pop Playlist Analysis:")
	fmt.Printf("Total pop songs: %d\n", len(popSet))
	fmt.Println("Songs in pop playlist:")
	for _, s := range getSortedKeys(popSet) {
		fmt.Printf("♪ %s\n", s)
	}

	// --- Combined ---
	combined := make(map[string]struct{})
	for s := range rockSet {
		combined[s] = struct{}{}
	}
	for s := range popSet {
		combined[s] = struct{}{}
	}

	fmt.Println("Combined Playlist Analysis:")
	fmt.Printf("Total unique songs: %d\n", len(combined))
	fmt.Println("All songs in combined playlist:")
	for _, s := range getSortedKeys(combined) {
		fmt.Printf("♫ %s\n", s)
	}

	// --- Summary ---
	fmt.Println("Playlist Summary:")
	fmt.Printf("Rock songs: %d\n", len(rockSet))
	fmt.Printf("Pop songs: %d\n", len(popSet))
	fmt.Printf("Combined unique songs: %d\n", len(combined))
}

func getSortedKeys(m map[string]struct{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
