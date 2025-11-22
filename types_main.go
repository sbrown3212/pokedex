package main

import "github.com/sbrown3212/pokedex/internal/pokeapi"

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

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
