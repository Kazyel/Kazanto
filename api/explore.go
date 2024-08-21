package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Kazyel/Poke-CLI/cache"
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
func UnmarshalExploreResponse(body []byte, exploreResponse *ExploreResponse) error {
	err := json.Unmarshal(body, &exploreResponse)

	if err != nil {
		fmt.Println("Error unmarshalling explore response.")
		return err
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
		fmt.Println("\n[CACHED] Exploring " + location + "...")

		err := UnmarshalExploreResponse(cachedExploreResponse, &exploreResponse)

		if err != nil {
			fmt.Println("Error unmarshalling explore response.")
		}

		fmt.Println("\nPokémons found:")
		for _, pokemon := range exploreResponse.PokemonEncounters {
			fmt.Println(pokemon.Pokemon.Name)
		}

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

	fmt.Println("\nExploring " + location + "...")

	response, err := http.Get(urlToSearch)

	if err != nil {
		fmt.Println("Error getting location.")
	}

	body, err := io.ReadAll(response.Body)
	exploreCache.AddToCache(urlToSearch, []byte(body))
	response.Body.Close()

	if response.StatusCode > 299 {
		fmt.Printf("Response failed with status code %d", response.StatusCode)
	}

	if err != nil {
		fmt.Println("Error reading response body.")
	}

	err = UnmarshalExploreResponse(body, &exploreResponse)

	if err != nil {
		fmt.Println("Error unmarshalling explore response.")
	}

	fmt.Println("\nPokémons found:")
	for _, pokemon := range exploreResponse.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
