package main

import (
	"time"

	"github.com/mikarwacki/pokedex/internal/pokeapi"
)

func main() {
	pokeApiClient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	config := &Config{pokeApiClient: pokeApiClient}
	startRepl(config)
}
