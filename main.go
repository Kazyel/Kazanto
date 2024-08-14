package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

		input := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch input {
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

		case "exit":
			commandMap["exit"].callback()

		default:
			fmt.Println("Invalid command")
		}
	}
}
