package api

import (
	"fmt"
	"math/rand"
	"time"
)

var pokedex *Pokedex = CreatePokedex()

func CatchPokemon(pokemon Pokemon) error {
	fmt.Println("Catching a Pokemon...")
	captureChance := rand.Intn(100)

	switch true {
	case captureChance >= 75 && captureChance <= 100:
		fmt.Println("You caught the Pokemon!")
		return pokedex.AddPokemon(pokemon.Name, pokemon.Type)

	case captureChance >= 25 && captureChance <= 75:
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second * 2)
			fmt.Println("Trying again...")

			captureRetry := rand.Intn(5)
			if captureRetry >= 4 {
				fmt.Println("You caught the Pokemon!")
				return nil
			}
		}

		fmt.Println("You failed to catch a Pokemon.")
		return nil

	case captureChance >= 0 && captureChance <= 25:
		fmt.Println("You failed to catch a Pokemon.")
		return nil

	default:
		fmt.Println("Something went wrong.")
		return nil
	}
}
