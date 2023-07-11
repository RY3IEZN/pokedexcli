package main

import "fmt"

func callbackHelp() error {
	fmt.Println("Welcome pokedex help menu!")
	fmt.Println("here are your available commands")

	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s \n", cmd.name, cmd.description)
	}

	fmt.Println("")
	return nil
}
