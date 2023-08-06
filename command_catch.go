package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callBackCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		fmt.Print(err)
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Print(pokemon.BaseExperience, randNum, threshold)
	if randNum < threshold {
		return fmt.Errorf("failed to capture %s, too strong", pokemon.Name)
	}

	fmt.Printf("a captured %s, nice!!! \n", pokemon.Name)
	return nil
}
