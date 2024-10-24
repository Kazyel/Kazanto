package api

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/Kazyel/Poke-CLI/utils"
)

type PokemonResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}
	}
	Types []struct {
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}
	}
	BaseExperience int `json:"base_experience"`
}

var pokemonResponse PokemonResponse = PokemonResponse{}

func (pokedex *Pokedex) CatchPokemon(pokemon string) error {
	if pokedex.Pokemons[pokemon].Name != "" {
		utils.PrintError("Pokemon already captured.")
		return fmt.Errorf("Pokemon already captured")
	}

	pokeRes, err := FetchPokemonData(pokemon)

	if err != nil {
		return err
	}

	utils.PrintAction("Catching "+pokeRes.Name, "primary")
	time.Sleep(time.Second * 1)

	expBase := pokeRes.BaseExperience
	captureChance := math.Floor(((rand.Float64() * 100) / float64(expBase)) * 100)

	switch true {
	case captureChance >= 75:
		utils.PrintSuccessfulCatch()
		return pokedex.AddPokemon(pokeRes.Name, pokeRes.Types[0].Type.Name)

	case captureChance >= 15 && captureChance <= 75:
		for i := 0; i < 2; i++ {
			time.Sleep(time.Second * 2)
			utils.PrintAction("Trying again", "secondary")

			captureRetry := math.Floor(((rand.Float64() * 100) / float64(expBase)) * 100)

			if captureRetry >= captureChance {
				utils.PrintSuccessfulCatch()
				return pokedex.AddPokemon(pokeRes.Name, pokeRes.Types[0].Type.Name)
			}
		}

		utils.PrintFailedCatch(pokeRes.Name)
		return nil

	case captureChance >= 0 && captureChance <= 15:
		utils.PrintFailedCatch(pokeRes.Name)
		return nil

	default:
		utils.PrintError("Something went wrong.")
		return nil
	}
}
