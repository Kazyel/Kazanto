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

var player *api.Player

func printWelcomeMessage() {
	fmt.Println("\n*-------------------------------*")
	fmt.Printf("\n Hey %s, welcome to Kazanto!\n\n", player.Name)
	fmt.Println(" Type 'help' for a list of commands.")
	fmt.Println(" Type 'exit' to exit.")
	fmt.Println("\n*-------------------------------*")
}

func printStartMessage() {
	fmt.Println("\n*-------------------------------*")
	fmt.Print("\n Welcome to Kazanto!\n\n")
	fmt.Println(" To start the game, please provide your name.")
	fmt.Println(" Type 'exit' to exit.")
	fmt.Println("\n*-------------------------------*")
}

func createPlayer() *api.Player {
	printStartMessage()

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

		if len(args) < 1 {
			utils.PrintError("Please provide your name.")
			continue
		}

		if len(args) > 1 {
			utils.PrintError("Too many arguments. Please provide only one name.")
			continue
		}

		if args[0] == "exit" {
			CommandExit()
		}

		player = api.CreateNewPlayer(args[0])
		break
	}

	return player
}

func gameLoop(commandMap map[int]cliCommand) {
	if player == nil {
		player = createPlayer()
	}

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
			if len(args) > 2 {
				utils.PrintError("Too many arguments.")
				continue
			}

			commandMap[0].callback()

			keySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
			for key := range keySlice {
				utils.PrintCommmands(commandMap[key].name, commandMap[key].description)
			}

		case "map":
			if len(args) > 2 {
				utils.PrintError("Too many arguments.")
				continue
			}

			commandMap[1].callback()

		case "mapback":
			if len(args) > 2 {
				utils.PrintError("Too many arguments.")
				continue
			}

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
			if len(args) < 3 {
				utils.PrintError("Please provide a pokémon to catch and a pokéball.")
				continue
			}

			if len(args) > 3 {
				utils.PrintError("Too many arguments. Please provide only one pokémon and one pokéball.")
				continue
			}

			ballType := strings.ToLower(args[2])
			ballTypes := []string{"pokeball", "masterball", "ultraball"}
			pokeballIsValid := false
			pokemonToCatch := strings.ToLower(args[1])

			for index := range ballTypes {
				if ballType == ballTypes[index] {
					pokeballs := player.Inventory.Pokeballs[ballType]
					pokeballIsValid = true

					if pokeballs.Quantity < 1 {
						utils.PrintError("You don't have any " + ballType + ".")
						break
					}

					player.Inventory.Pokeballs[args[2]].Quantity--
					commandMap[4].callback(pokemonToCatch, player, pokeballs.Pokeball)
					break
				}
			}

			if !pokeballIsValid {
				utils.PrintError("Please provide a valid pokéball.")
				continue
			}

		case "inspect":
			if len(args) < 2 {
				utils.PrintError("Please provide only one pokémon to inspect.")
				continue
			}

			if len(args) > 2 {
				utils.PrintError("Too many arguments. Please provide only one pokémon.")
				continue
			}

			commandMap[5].callback(args[1], player.Pokedex)

		case "pokedex":
			if len(args) > 2 {
				utils.PrintError("Too many arguments.")
				continue
			}

			commandMap[6].callback(player.Pokedex)

		case "inventory":
			if len(args) > 2 {
				utils.PrintError("Too many arguments.")
				continue
			}

			commandMap[7].callback(player)

		case "exit":
			commandMap[8].callback()

		default:
			utils.PrintError("Invalid command.")
		}
	}
}

func main() {
	commandMap := Commands()
	gameLoop(commandMap)
}
