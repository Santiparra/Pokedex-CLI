package main

import(
	"time"
	"github.com/Santiparra/Pokedex-CLI/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	runRepl(cfg)
}