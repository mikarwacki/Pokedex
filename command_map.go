package main

import (
	"fmt"
)

func commandMap(config *Config, args ...string) error {
	location, err := config.pokeApiClient.ListLocations(config.NextURL)
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
