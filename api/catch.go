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
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	BaseExperience int `json:"base_experience"`
	Height         int `json:"height"`
	Weight         int `json:"weight"`
}

var pokemonResponse PokemonResponse = PokemonResponse{}

func generateCatchChance(expBase int) int {
	captureChance := math.Floor(((rand.Float64() * 100) / float64(expBase)) * 100)
	return int(captureChance)
}

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
	captureChance := generateCatchChance(expBase)

	switch true {
	case captureChance >= 75:
		utils.PrintSuccessfulCatch()
		return pokedex.AddPokemon(pokeRes)

	case captureChance >= 15 && captureChance <= 75:
		for i := 0; i < 2; i++ {
			utils.PrintAction("Trying again", "secondary")
			time.Sleep(time.Second * 2)

			captureRetry := generateCatchChance(expBase)

			if captureRetry >= captureChance {
				utils.PrintSuccessfulCatch()
				return pokedex.AddPokemon(pokeRes)
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
