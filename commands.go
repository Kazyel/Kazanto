package main

import (
	"fmt"
	"os"

	"github.com/Kazyel/Poke-CLI/api"
	"github.com/Kazyel/Poke-CLI/utils"
)

type cliCommand struct {
	name        string
	description string
	callback    func(...interface{}) error
}

func CommandExit() error {
	fmt.Println("\nExiting Poke-CLI...")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	utils.PrintTitle("These are the available commands:")
	return nil
}

// Commands returns a map of commands.
func Commands() map[int]cliCommand {
	return map[int]cliCommand{
		0: {
			name:        "help",
			description: "Displays a help message.",
			callback:    func(...interface{}) error { return commandHelp() },
		},
		1: {
			name:        "map",
			description: "Displays the next 20 maps.",
			callback:    func(...interface{}) error { return api.GetNextLocations() },
		},
		2: {
			name:        "mapback",
			description: "Displays the previous 20 maps.",
			callback:    func(...interface{}) error { return api.GetPreviousLocations() },
		},
		3: {
			name:        "explore",
			description: "Explore a location.",
			callback: func(location ...interface{}) error {
				return api.ExploreLocation(location[0].(string))
			},
		},
		4: {
			name:        "catch",
			description: "Catch a Pokemon.",
			callback: func(args ...interface{}) error {
				return args[1].(*api.Player).
					Pokedex.CatchPokemon(args[0].(string), args[2].(api.Pokeball))
			},
		},
		5: {
			name:        "inspect",
			description: "Inspect a Pokemon.",
			callback: func(pokemon ...interface{}) error {
				pokemon[1].(*api.Pokedex).InspectPokemon(pokemon[0].(string))
				return nil
			},
		},
		6: {
			name:        "pokedex",
			description: "Displays the Pokedex.",
			callback: func(pokedex ...interface{}) error {
				pokedex[0].(*api.Pokedex).RenderPokedex()
				return nil
			},
		},
		7: {
			name:        "inventory",
			description: "Displays the player inventory.",
			callback: func(player ...interface{}) error {
				player[0].(*api.Player).InspectInventory()
				return nil
			},
		},
		8: {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    func(...interface{}) error { return CommandExit() },
		},
	}
}
