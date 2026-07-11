package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonResp, err := cfg.pokeapiClient.PokemonStats(args[0])
	if err != nil {
		return err
	}

	pokeName := pokemonResp.Name
	catchChance := rand.IntN(pokemonResp.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)
	if catchChance > 40 {
		fmt.Printf("%s was caught!\n", pokeName)
		cfg.caughtPokemon[pokeName] = pokemonResp
	} else {
		fmt.Printf("%s escaped!\n", pokeName)
	}
	return nil
}
