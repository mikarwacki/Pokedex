package main

import (
	"fmt"
)

func commandMapB(config *Config) error {
	location, err := config.pokeApiClient.ListLocations(config.PreviousURL)
	if err != nil {
		fmt.Println("Error getting the locations")
		return err
	}

	for _, result := range location.Results {
		fmt.Println(result.Name)
	}

	config.PreviousURL = location.Previous
	config.NextURL = location.Next
	return nil
}
