package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapB(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("You're at the beginning of the list")
		return nil
	}

	res, err := http.Get(cfg.Previous)
	if err != nil {
		return fmt.Errorf("'mapb' GET request error: %s", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("'mapb' io ReadAll error: %s", err)
	}

	var locations locationAreaResponse
	if err := json.Unmarshal(data, &locations); err != nil {
		return fmt.Errorf("'mapb' unmarshal error: %s", err)
	}

	if locations.Next == nil {
		cfg.Next = ""
	} else {
		cfg.Next = *locations.Next
	}
	if locations.Previous == nil {
		cfg.Previous = ""
	} else {
		cfg.Previous = *locations.Previous
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
