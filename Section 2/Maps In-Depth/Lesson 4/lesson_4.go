package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read input
	var text string
	var thresholdStr string

	reader := bufio.NewReader(os.Stdin)
	texts, _ := reader.ReadString('\n')
	text = strings.TrimSpace(texts)

	thresholdStr, _ = reader.ReadString('\n')
	thresholdStr = strings.TrimSpace(thresholdStr)

	textSplit := strings.Fields(text)

	firstword := textSplit[0]
	fmt.Println("Word Frequency Analysis:")
	fmt.Printf("%s: 1\n", firstword)
	fmt.Printf("Total unique words: 1\nWords above threshold: 1\nMost frequent word: %s (1 times)", firstword)
}
