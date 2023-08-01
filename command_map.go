package main

import (
	"errors"
	"fmt"
)

func callBackMap(cfg *config, args ...string) error {

	resp, err := cfg.pokeapiClient.ListLocationArea(cfg.nextLocationAreaUrl)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Available Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s \n", area.Name)

	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}

func callBackMapb(cfg *config, args ...string) error {

	if cfg.prevLocationAreaUrl == nil {
		return errors.New("you are on the firstpage")
	}
	resp, err := cfg.pokeapiClient.ListLocationArea(cfg.prevLocationAreaUrl)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Available Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s \n", area.Name)

	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}
