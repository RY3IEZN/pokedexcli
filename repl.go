package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// intiate the scanner and get the user input
func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// show a pointer to indicate to type
		fmt.Print(" Pokedexcli~> ")

		// intiate a scanner,get the input and point it to text
		scanner.Scan()
		text := scanner.Text()

		// clean the input conv to lowercase
		cleaned := clearInput(text)

		// if the input is empty, loopback to the top
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}

		err := command.callBack(cfg)
		if err != nil {
			fmt.Print(err)
		}

		// print the word that was inputed
		// fmt.Println("echo,", cleaned)
	}

}

type cliCommand struct {
	name        string
	description string
	callBack    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "prints the help menu",
			callBack:    callbackHelp,
		},

		"exit": {
			name:        "exit",
			description: "stops the pokedexcli",
			callBack:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "list available location",
			callBack:    callBackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "list previous locations",
			callBack:    callBackMapb,
		},
	}
}

// clear the user input
func clearInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
