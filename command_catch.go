package main

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"slices"
)

// TODO: move to types main file

type Pokemon struct {
	id             int
	name           string
	baseExperience int
	height         int
	weight         int
	stats          []struct {
		name string
		val  int
	}
	types []struct {
		name string
		slot int
	}
}

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

	var stats []struct {
		name string
		val  int
	}
	for _, stat := range pokemonResp.Stats {
		newStat := struct {
			name string
			val  int
		}{
			name: stat.Stat.Name,
			val:  stat.BaseStat,
		}
		stats = append(stats, newStat)
	}

	var types []struct {
		name string
		slot int
	}
	for _, respPokeType := range pokemonResp.Types {
		newPokeType := struct {
			name string
			slot int
		}{
			name: respPokeType.Type.Name,
			slot: respPokeType.Slot,
		}
		types = append(types, newPokeType)
	}

	slices.SortFunc(types, func(a, b struct {
		name string
		slot int
	},
	) int {
		return cmp.Compare(a.slot, b.slot)
	})

	pokemon := Pokemon{
		id:             pokemonResp.ID,
		name:           pokemonResp.Name,
		baseExperience: pokemonResp.BaseExperience,
		height:         pokemonResp.Height,
		weight:         pokemonResp.Weight,
		stats:          stats,
		types:          types,
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
