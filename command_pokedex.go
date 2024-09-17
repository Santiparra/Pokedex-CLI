package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	pokemonList := cfg.caughtPokemon
	if len(pokemonList) < 1 {
		fmt.Println("You havent caught any pokemon yet, go ahead, explore zones and catch some!")
	}

	fmt.Println("Pokedex:")
	for _, pokemon := range pokemonList {
		fmt.Println("  -", pokemon.Name)
	}

	return nil
}
