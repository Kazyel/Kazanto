package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Kazyel/Poke-CLI/cache"
	"github.com/Kazyel/Poke-CLI/utils"
)

type LocationResponse struct {
	Next     any `json:"next"`
	Previous any `json:"previous"`

	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var locations LocationResponse = LocationResponse{}
var locationsCache *cache.Cache = cache.NewCache(300 * time.Second)

// unmarshalLocations unmarshals the response from the PokeAPI.
func unmarshalLocations(body []byte, locations *LocationResponse) error {
	err := json.Unmarshal(body, &locations)

	if err != nil {
		utils.PrintError("Error unmarshalling locations.")
		return fmt.Errorf("error unmarshalling locations")
	}

	return nil
}

// printLocations prints the locations from the PokeAPI.
func printLocations(locations LocationResponse) {
	utils.PrintTitle("Current locations:")

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
}

/*
checkLocationCache checks if the response from the PokeAPI is cached.
If it is, it prints the cached response and returns true.
Otherwise, it returns false.
*/
func checkLocationCache(url string) bool {
	cachedLocations, ok := locationsCache.GetFromCache(url)

	if ok {
		err := unmarshalLocations(cachedLocations, &locations)

		if err != nil {
			log.Fatalf("Error unmarshalling locations response.")
		}

		if locations.Results != nil {
			utils.PrintCachedAction("Getting next locations")
			printLocations(locations)
		}

		return true
	}

	return false
}

/*
GetNextLocations sends a request to the PokeAPI to get the next 20 locations.
It first checks if the response is cached, and if it is, it prints the cached response.
Otherwise, it sends a request to the PokeAPI and prints the response.
*/
func GetNextLocations() error {
	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	if locations.Next == nil && locations.Results != nil {
		utils.PrintError("No more next locations.")
		return nil
	}

	if locations.Next != nil {
		url = locations.Next.(string)
	}

	cacheExists := checkLocationCache(url)

	if cacheExists {
		return nil
	}

	response, err := http.Get(url)

	if err != nil {
		utils.PrintError("Error getting locations.")
		return fmt.Errorf("error getting locations")
	}

	body, err := io.ReadAll(response.Body)
	locationsCache.AddToCache(url, []byte(body))
	response.Body.Close()

	if response.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d", response.StatusCode)
	}

	if err != nil {
		log.Fatalf("Error reading response body.")
	}

	err = unmarshalLocations(body, &locations)

	if err != nil {
		log.Fatalf("Error unmarshalling locations response.")
	}

	if locations.Results != nil {
		utils.PrintAction("Getting next locations", "primary")
		printLocations(locations)
	}

	return nil
}

/*
GetPreviousLocations sends a request to the PokeAPI to get the previous 20 locations.
It first checks if the response is cached, and if it is, it prints the cached response.
Otherwise, it sends a request to the PokeAPI and prints the response.
*/
func GetPreviousLocations() error {
	var url string

	if locations.Previous == nil {
		utils.PrintError("No previous locations.")
		locations.Next = "https://pokeapi.co/api/v2/location-area/"
		return nil
	}

	url = locations.Previous.(string)
	cachedLocations, ok := locationsCache.GetFromCache(url)

	if ok {
		err := unmarshalLocations(cachedLocations, &locations)

		if err != nil {
			log.Fatalf("Error unmarshalling locations response.")
		}

		if locations.Results != nil {
			utils.PrintCachedAction("Getting previous locations")
			printLocations(locations)
		}

		return nil
	}

	response, err := http.Get(url)

	if err != nil {
		utils.PrintError("Error getting locations.")
		return fmt.Errorf("error getting locations")
	}

	body, err := io.ReadAll(response.Body)
	locationsCache.AddToCache(url, []byte(body))
	response.Body.Close()

	if response.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d", response.StatusCode)
	}

	if err != nil {
		log.Fatalf("Error reading response body.")
	}

	err = unmarshalLocations(body, &locations)

	if err != nil {
		log.Fatalf("Error unmarshalling locations response.")
	}

	if locations.Results != nil {
		utils.PrintAction("Getting previous locations", "primary")
		printLocations(locations)
	}

	return nil
}
