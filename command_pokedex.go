package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemons := range cfg.caughtPokemon {
		fmt.Println(" - ", pokemons.Name)
	}
	return nil
}
