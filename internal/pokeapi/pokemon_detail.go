package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonDetail(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon" + "/" + name

	val, ok := c.cache.Get(url)
	if ok {
		var cachedPokemon RespPokemon
		err := json.Unmarshal(val, &cachedPokemon)
		if err != nil {
			return RespPokemon{}, err
		}

		return cachedPokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return RespPokemon{}, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, data)

	var pokemonDetail RespPokemon
	err = json.Unmarshal(data, &pokemonDetail)
	if err != nil {
		return RespPokemon{}, err
	}

	return pokemonDetail, nil
}
