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

		fmt.Printf("Your command was: %v\n", commandName)
	}
}
