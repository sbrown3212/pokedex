package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.name)
	fmt.Printf("Height: %d\n", pokemon.height)
	fmt.Printf("Weight: %d\n", pokemon.weight)
	fmt.Println("Stats:")
	for _, pokeStat := range pokemon.stats {
		fmt.Printf(" - %s: %d\n", pokeStat.name, pokeStat.val)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.types {
		fmt.Printf(" - %s\n", pokeType.name)
	}

	return nil
}
