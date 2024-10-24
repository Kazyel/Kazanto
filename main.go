package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kazyel/Poke-CLI/api"
	"github.com/Kazyel/Poke-CLI/utils"
	"github.com/fatih/color"
)

func printWelcomeMessage() {
	fmt.Println("\n*-------------------------------*")
	fmt.Println("\n Hey, welcome to Poke-CLI!")
	fmt.Println(" Type 'help' for a list of commands.")
	fmt.Println(" Type 'exit' to exit.")
	fmt.Println("\n*-------------------------------*")
}

func main() {
	var pokedex *api.Pokedex = api.CreatePokedex()
	commandMap := Commands()
	printWelcomeMessage()

	for {
		color.New(color.FgHiCyan, color.Bold).Fprint(os.Stdout, "\nPoke-CLI > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		err := scanner.Err()
		if err != nil {
			utils.PrintError("Error reading input.")
			continue
		}

		commandLine := strings.ToLower(strings.TrimSpace(scanner.Text()))
		args := strings.Split(commandLine, " ")

		switch args[0] {
		case "help":
			commandMap[0].callback()

			keySlice := []int{1, 2, 3, 4, 5, 6}
			for key := range keySlice {
				utils.PrintCommmands(commandMap[key].name, commandMap[key].description)
			}

		case "map":
			commandMap[1].callback()

		case "mapback":
			commandMap[2].callback()

		case "explore":
			if len(args) < 2 {
				utils.PrintError("Please provide one location.")
				continue
			}

			if len(args) > 2 {
				utils.PrintError("Too many arguments. Please provide only one location.")
				continue
			}

			commandMap[3].callback(args[1])

		case "catch":
			commandMap[4].callback(args[1], pokedex)

		case "pokedex":
			commandMap[5].callback(pokedex)

		case "exit":
			commandMap[6].callback()

		default:
			utils.PrintError("Invalid command.")
		}
	}
}
