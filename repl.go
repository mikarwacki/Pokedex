package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := Config{NextURL: locationUrl}
	for {
		fmt.Print("Pokedex> ")
		scanner.Scan()
		userInput := scanner.Text()
		commands := getCommands()
		words := clearInput(userInput)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if command, ok := commands[commandName]; ok {
			err := command.callback(&config)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
}

func clearInput(input string) []string {
	clear := strings.ToLower(input)
	words := strings.Fields(clear)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Fetches next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Fetches previous 20 locations",
			callback:    commandMapB,
		},
	}
}

type Location struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	NextURL     string
	PreviousURL *string
}
