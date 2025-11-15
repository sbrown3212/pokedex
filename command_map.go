package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *config) error {
	const locationAreaURL string = "https://pokeapi.co/api/v2/location-area"

	// Initialize config
	if cfg.Next == "" && cfg.Previous == "" {
		cfg.Next = locationAreaURL
	}

	if cfg.Next == "" {
		fmt.Println("You've reached the end of the list")
		return nil
	}

	res, err := http.Get(cfg.Next)
	if err != nil {
		return fmt.Errorf("location area GET request error: %s", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("location area io ReadAll error: %s", err)
	}

	var locations locationAreaResponse
	if err := json.Unmarshal(data, &locations); err != nil {
		return fmt.Errorf("location area unmarshal error: %s", err)
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

type locationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
