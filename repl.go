package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	prompt string = "Pokedex > "
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, words[1:]...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	words := strings.Fields(lowercase)

	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <location-name>",
			description: "List the pokemon of a given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Attepmt to catch a given pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "explore <pokemon-name>",
			description: "View details of pokemon you have caught",
			callback:    commandInspect,
		},
	}
}
