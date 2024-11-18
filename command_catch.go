package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide pokemon name ")
	}

	name := args[0]
	pokemon, err := config.pokeApiClient.CatchPokemon(name)
	if err != nil {
		return err
	}

	num := rand.IntN(pokemon.BaseExperience)
	if num > 19 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	config.pokedex[name] = pokemon
	fmt.Printf("%s was caught\n", name)
	return nil
}
