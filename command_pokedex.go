package main

import "fmt"

func commandPokedex(cfg *config, _ ...string) error {
	pokedex := cfg.pokedex

	fmt.Println("Your pokemon:")
	for _, pokemon := range pokedex {
		fmt.Printf(" - %s\n", pokemon.name)
	}
	return nil
}
