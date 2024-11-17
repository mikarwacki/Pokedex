package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapB(config *Config) error {
	if (config.PreviousURL) == nil {
		fmt.Print("There is no previous location")
		return nil
	}

	res, err := http.Get(*(config.PreviousURL))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var location Location
	err = json.Unmarshal(data, &location)
	if err != nil {
		return err
	}
	for _, result := range location.Results {
		fmt.Println(result.Name)
	}

	config.PreviousURL = location.Previous
	config.NextURL = location.Next
	return nil
}
