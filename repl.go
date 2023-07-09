package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// intiate the scanner and get the user input
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")

		scanner.Scan()
		text := scanner.Text()

		cleaned := clearInput(text)
		if len(cleaned) == 0 {
			continue
		}

		fmt.Println("echoh,", cleaned)
	}

}

// clear the user input
func clearInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
