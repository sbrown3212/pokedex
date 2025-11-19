package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		// return errors.New("you must provide a location name")
		return fmt.Errorf("you must provide a location name")
	}

	area := args[0]

	areaDetail, err := cfg.pokeapiClient.LocationDetail(area)
	if err != nil {
		return fmt.Errorf("unable to fetch location detail: %s", err)
	}

	for _, encounter := range areaDetail.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
