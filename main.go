package main

import(
	"time"
	"github.com/Santiparra/Pokedex-CLI/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}
	runRepl(cfg)
}