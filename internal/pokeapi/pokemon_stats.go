package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonStats(pokemonName string) (Pokemon, error) {
	url := pokemonURL + pokemonName

	if value, ok := c.cache.Get(url); ok {
		pokeResp := Pokemon{}
		err := json.Unmarshal(value, &pokeResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokeResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	pokeResp := Pokemon{}
	err = json.Unmarshal(data, &pokeResp)
	if err != nil {
		return Pokemon{}, err
	}

	return pokeResp, nil
}
