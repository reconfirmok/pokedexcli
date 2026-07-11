package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	if p, ok := cfg.caughtPokemon[args[0]]; ok {
		fmt.Printf("Name: %s\n", p.Name)
		fmt.Printf("Height: %d\n", p.Height)
		fmt.Printf("Weight: %d\n", p.Weight)
		fmt.Println("Stats:")
		for _, s := range p.Stats {
			fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range p.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
