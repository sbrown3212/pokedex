package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonResp, err := cfg.pokeapiClient.PokemonDetail(pokemonName)
	if err != nil {
		return fmt.Errorf("unable to fetch pokemon details: %s", err)
	}

	pokemon := Pokemon{
		id:             pokemonResp.ID,
		name:           pokemonResp.Name,
		baseExperience: pokemonResp.BaseExperience,
	}

	rollResult := rand.Float64()
	target := 1.9/(1.0+0.004*float64(pokemon.baseExperience)) - 0.7

	caught := rollResult < target

	if caught {
		fmt.Printf("%s was caught!\n", pokemon.name)
		cfg.pokedex[pokemon.name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.name)
	}

	return nil
}

type Pokemon struct {
	id             int
	name           string
	baseExperience int
}
