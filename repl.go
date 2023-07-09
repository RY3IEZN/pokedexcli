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
		// show a pointer to indicate to type
		fmt.Print(" > ")

		// intiate a scanner,get the input and point it to text
		scanner.Scan()
		text := scanner.Text()

		// clean the input conv to lowercase
		cleaned := clearInput(text)

		// if the input is empty, loopback to the top
		if len(cleaned) == 0 {
			continue
		}

		// switch to check the command that was inputed
		command := cleaned[0]
		switch command {
		case "help":
			fmt.Println("Welcome pokedex help menu!")
			fmt.Println("here are your available commands")
			fmt.Println(" - help")
			fmt.Println(" - exit")
			fmt.Println("")
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid Command")
		}

		// print the word that was inputed
		fmt.Println("echo,", cleaned)
	}

}

// clear the user input
func clearInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
