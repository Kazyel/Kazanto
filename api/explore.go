package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Kazyel/Poke-CLI/cache"
	"github.com/Kazyel/Poke-CLI/utils"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type ExploreResponse struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

var exploreResponse ExploreResponse = ExploreResponse{}
var exploreCache *cache.Cache = cache.NewCache(300 * time.Second)

// UnmarshalExploreResponse unmarshals the response from the PokeAPI.
func UnmarshalResponse(body []byte, response interface{}) error {
	err := json.Unmarshal(body, &response)

	if err != nil {
		utils.PrintError("Error unmarshalling response.")
		return fmt.Errorf("error unmarshalling response")
	}

	return nil
}

/*
checkExploreCache checks if the response from the PokeAPI is cached.
If it is, it prints the cached response and returns true.
Otherwise, it returns false.
*/
func checkExploreCache(url string, location string) bool {
	cachedExploreResponse, ok := exploreCache.GetFromCache(url)

	if ok {
		utils.PrintCachedAction("Exploring " + location)
		err := UnmarshalResponse(cachedExploreResponse, &exploreResponse)

		if err != nil {
			utils.PrintError("Error unmarshalling explore response.")
			return false
		}

		utils.PrintTitle("Pokémons found:")

		headerFmt := color.New(color.FgHiBlue, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("Name", "Type")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for _, pokemon := range exploreResponse.PokemonEncounters {
			pokemonResponse, err := FetchPokemonData(pokemon.Pokemon.Name)

			if err != nil {
				utils.PrintError("Error fetching pokemon data.")
				return false
			}

			tbl.AddRow(pokemonResponse.Name, pokemonResponse.Types[0].Type.Name)
		}

		tbl.Print()
		return true
	}

	return false
}

/*
ExploreLocation sends a request to the PokeAPI to explore a location.
It first checks if the response is cached, and if it is, it prints the cached response.
Otherwise, it sends a request to the PokeAPI and prints the response.
*/
func ExploreLocation(location string) error {
	urlToSearch := "https://pokeapi.co/api/v2/location-area/" + location
	cacheExists := checkExploreCache(urlToSearch, location)

	if cacheExists {
		return nil
	}

	utils.PrintAction("Exploring "+location, "primary")

	response, err := http.Get(urlToSearch)

	if response.StatusCode > 299 {
		utils.PrintError("No data found for location.")
		return fmt.Errorf("no data found for location")
	}

	if err != nil {
		utils.PrintError("Error fetching location data.")
		return fmt.Errorf("error fetching location data")
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	exploreCache.AddToCache(urlToSearch, []byte(body))

	if err != nil {
		utils.PrintError("Error reading response body.")
		return fmt.Errorf("error reading response body")
	}

	err = UnmarshalResponse(body, &exploreResponse)

	if err != nil {
		utils.PrintError("Error unmarshalling explore response.")
		return fmt.Errorf("error unmarshalling explore response")
	}

	utils.PrintTitle("Pokémons found:")
	headerFmt := color.New(color.FgHiBlue, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Name", "Type")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, pokemon := range exploreResponse.PokemonEncounters {
		pokemonResponse, err := FetchPokemonData(pokemon.Pokemon.Name)

		if err != nil {
			utils.PrintError("Error fetching pokemon data.")
			return fmt.Errorf("error fetching pokemon data")
		}

		tbl.AddRow(pokemonResponse.Name, pokemonResponse.Types[0].Type.Name)
	}

	tbl.Print()
	return nil
}
