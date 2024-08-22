package main

import (
	"fmt"
	"os"

	"github.com/Kazyel/Poke-CLI/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(...interface{}) error
}

func commandExit() error {
	fmt.Println("\nExiting Poke-CLI...")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("\nThese are the available commands:\n\n")
	return nil
}

// Commands returns a map of commands.
func Commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    func(...interface{}) error { return commandHelp() },
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 maps.",
			callback:    func(...interface{}) error { return api.GetNextLocations() },
		},
		"mapback": {
			name:        "mapback",
			description: "Displays the previous 20 maps.",
			callback:    func(...interface{}) error { return api.GetPreviousLocations() },
		},
		"explore": {
			name:        "explore",
			description: "Explore a location.",
			callback: func(location ...interface{}) error {
				return api.ExploreLocation(location[0].(string))
			},
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon.",
			callback:    func(pokemon ...interface{}) error { return api.CatchPokemon(pokemon[0].(api.Pokemon)) },
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    func(...interface{}) error { return commandExit() },
		},
	}
}
