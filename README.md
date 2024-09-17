# Pokedex-CLI

This is a command-line interface (CLI) tool built using Go. It interacts with the [PokéAPI](https://pokeapi.co/) to provide Pokémon-related functionalities like exploring locations, catching Pokémon, and inspecting details about captured Pokémon.

This project is a learning exercise. Therefore, precompiled binaries are not included in the repository. Users interested in trying the tool can follow the installation and usage instructions below.

## Features

- **Explore Locations**: Navigate through different Pokémon locations.
- **Catch Pokémon**: Attempt to catch wild Pokémon.
- **Inspect Captured Pokémon**: View details about the Pokémon you have caught.
- **View Your Pokédex**: Get a list of all captured Pokémon for further inspection.

## Commands

| Command                   | Description                                                      |
|----------------------------|------------------------------------------------------------------|
| `help`                     | Display help message                                             |
| `exit`                     | Exit the CLI                                                     |
| `map`                      | Go to the next page of locations                                 |
| `mapb`                     | Go to the previous page of locations                             |
| `explore <location_name>`   | Explore a specific location                                      |
| `catch <pokemon_name>`      | Attempt to catch a Pokémon                                       |
| `inspect <pokemon_name>`    | Inspect details of any captured Pokémon                           |
| `pokedex`                  | Get a list of all captured Pokémon available for inspection      |

## Installation

To try out the Pokémon CLI Tool, you'll need to have [Go](https://golang.org/dl/) installed on your system. You can install it by following the instructions on the official Go website.

Once Go is installed, you can clone this repository and build the tool locally.

### Steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/Santiparra/Pokedex-CLI.git
   ```
2. Move to the root project folder just created:   
   ```bash
   cd Pokedex-CLI
   ```
3. Build the CLI tool:
   ```bash
   go build -o pokemon-cli
   ```
### Run the tool:
#### On Windows:
   ```bash
   pokemon-cli.exe
   ```
#### On Linux/macOS:
   ```bash
   ./pokemon-cli
   ```
## How to Use

After running the tool, you can enter any of the commands listed above to interact with the PokéAPI. 

## PokéAPI

This CLI tool uses the [PokéAPI](https://pokeapi.co/)  to fetch Pokémon data. If you're interested in the API itself or want to explore further, you can visit their website.