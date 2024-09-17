package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/Santiparra/Pokedex-CLI/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func runRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("PokeDex >")
		reader.Scan()

		inputs := cleanInput(reader.Text())
		if len(inputs) == 0 {
			continue
		}

		commandName := inputs[0]

		command, exist := getCommands()[commandName]
		if exist {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Command not Found")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Display help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit CLI",
			callback: commandExit,
		},
		"map": {
			name:        "map",
			description: "Go to locations next page",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Go to locations previous page",
			callback:    commandMapb,
		},
	}
}