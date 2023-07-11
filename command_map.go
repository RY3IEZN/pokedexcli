package main

import (
	"fmt"
	"log"
	"pokedexcli/internal/pokeapi"
)

func callBackMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationArea()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Available Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s \n", area.Name)

	}
	return nil
}
