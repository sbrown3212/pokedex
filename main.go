package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	prompt string = "Pokedex > "
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		fmt.Printf("Your command was: %v\n", cleaned[0])
	}
}
