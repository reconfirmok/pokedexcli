package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(areaName string) (AreaPokemonEncounter, error) {
	url := locationURL + areaName

	if value, ok := c.cache.Get(url); ok {
		areaResp := AreaPokemonEncounter{}
		err := json.Unmarshal(value, &areaResp)
		if err != nil {
			return AreaPokemonEncounter{}, err
		}
		return areaResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaPokemonEncounter{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaPokemonEncounter{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaPokemonEncounter{}, err
	}

	c.cache.Add(url, data)

	areaResp := AreaPokemonEncounter{}
	err = json.Unmarshal(data, &areaResp)
	if err != nil {
		return AreaPokemonEncounter{}, err
	}

	return areaResp, nil
}
