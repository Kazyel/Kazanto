package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kazyel/Poke-CLI/api"
)

func printWelcomeMessage() {
	fmt.Println("\n*-------------------------------*")
	fmt.Println("\n Hey, welcome to Poke-CLI!")
	fmt.Println(" Type 'help' for a list of commands.")
	fmt.Println(" Type 'exit' to exit.")
	fmt.Println("\n*-------------------------------*")
}

func main() {
	commandMap := Commands()
	printWelcomeMessage()

	for {
		fmt.Print("\nPoke-CLI > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		err := scanner.Err()
		if err != nil {
			fmt.Println("Error reading input")
			continue
		}

		commandLine := strings.ToLower(strings.TrimSpace(scanner.Text()))
		args := strings.Split(commandLine, " ")

		switch args[0] {
		case "help":
			commandMap["help"].callback()

			for _, command := range commandMap {
				fmt.Print(command.name + ": ")
				fmt.Println(command.description)
			}

		case "map":
			commandMap["map"].callback()

		case "mapback":
			commandMap["mapback"].callback()

		case "explore":
			if len(args) < 2 {
				fmt.Println("Please provide a location.")
				continue
			}

			if len(args) > 2 {
				fmt.Println("Too many arguments. Please provide only one location.")
				continue
			}

			commandMap["explore"].callback(args[1])

		case "catch":
			pokemon := api.Pokemon{
				Name: "Pikachu2",
				Type: "Electric",
			}
			commandMap["catch"].callback(pokemon)

		case "exit":
			commandMap["exit"].callback()

		default:
			fmt.Println("Invalid command")
		}
	}
}
