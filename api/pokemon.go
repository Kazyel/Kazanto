package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Kazyel/Poke-CLI/utils"
)

func FetchPokemonData(pokemon string) (PokemonResponse, error) {
	urlToSearch := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	response, err := http.Get(urlToSearch)

	if response.StatusCode > 299 {
		utils.PrintError("No data for pokémon.")
		return PokemonResponse{}, fmt.Errorf("no data for pokemon")
	}

	if err != nil {
		utils.PrintError("Error fetching pokemon data.")
		return PokemonResponse{}, fmt.Errorf("error fetching pokemon data")
	}
	body, err := io.ReadAll(response.Body)

	if err != nil {
		utils.PrintError("Error reading response body.")
		return PokemonResponse{}, fmt.Errorf("error reading response body")
	}

	defer response.Body.Close()
	err = UnmarshalResponse(body, &pokemonResponse)

	if err != nil {
		utils.PrintError("Error unmarshalling pokemon response.")
		return PokemonResponse{}, fmt.Errorf("error unmarshalling pokemon response")
	}

	return pokemonResponse, nil
}
