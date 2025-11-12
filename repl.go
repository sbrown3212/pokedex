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

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	words := strings.Fields(lowercase)

	return words
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		switch commandName {
		case "exit":
			err := commands["exit"].callback()
			if err != nil {
				fmt.Printf("error exiting Pokedex: %v\n", err)
			}
		case "help":
			err := commands["help"].callback()
			if err != nil {
				fmt.Printf("error with help command: %v\n", err)
			}
		default:
			fmt.Println("Unknown command")
		}

	}
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf(`Welcome to the Pokedex!
	Usage:

	help: Displays a help message
	exit: Exit the Pokedex`)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
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
}
