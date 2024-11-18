package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide pokemon name ")
	}

	name := args[0]
	if pokemon, ok := config.pokedex[name]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)

		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("  - %v\n", t.Type.Name)
		}

	} else {
		fmt.Printf("You have not caught %s yet\n", name)
	}

	return nil
}
