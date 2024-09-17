package main

import(
	"time"
	"github.com/Santiparra/Pokedex-CLI/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}
	runRepl(cfg)
}