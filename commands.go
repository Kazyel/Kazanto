package main

import (
	"fmt"
	"os"

	"github.com/Kazyel/Poke-CLI/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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

func Commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 maps.", 
			callback:    api.GetNextLocations,
		},
		"mapback": {
			name:        "mapback",
			description: "Displays the previous 20 maps.", 
			callback:    api.GetPreviousLocations,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
	}
}
