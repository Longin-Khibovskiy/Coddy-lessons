package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numExistingTasks := scanner.Text()
	num, _ := strconv.Atoi(strings.TrimSpace(numExistingTasks))
	scanner.Scan()
	scanner.Text()
	fmt.Printf("Total: %d tasks (0 completed, %d remaining)", num, num)
}
