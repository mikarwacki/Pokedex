package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide location name")
	}

	locationDetails, err := config.pokeApiClient.ExploreLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationDetails.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
