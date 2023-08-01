package main

import (
	"errors"
	"fmt"
)

func callBackExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)

	if err != nil {
		fmt.Print(err)
		fmt.Print("eoor")
	}

	fmt.Printf("Pokemons in %s:\n", locationArea.Name)

	for _, Pokemons := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s \n", Pokemons.Pokemon.Name)

	}
	return nil
}
