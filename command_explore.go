package main

import "fmt"

func commandExplore(config *Config, commandParam string) error {
	locationDetails, err := config.pokeApiClient.ExploreLocation(commandParam)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", commandParam)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationDetails.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
