package main

import "fmt"

func commandPokedex(config *Config, args ...string) error {
	fmt.Println("Your pokedex:")
	for name := range config.pokedex {
		fmt.Printf("  - %s\n", name)
	}
	return nil
}
