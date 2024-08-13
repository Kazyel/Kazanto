package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func unmarshalLocations(body []byte, locations *LocationResponse) error {
	err := json.Unmarshal(body, &locations)

	if err != nil {
		fmt.Println("Error unmarshalling locations.")
		return err
	}

	return nil
}

func printLocations(locations LocationResponse) {
	fmt.Print("Current locations:\n\n")

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
}

func GetNextLocations() error {
	var url string

	if locations.Next == nil && locations.Results != nil {
		fmt.Println("No more next locations.")
		return nil
	} else if locations.Next != nil {
		url = locations.Next.(string)
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error getting next locations.")
	}

	body, err := io.ReadAll(response.Body)
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

	if locations.Results == nil {
		fmt.Println("\nNo more next locations.")
		return nil
	}

	if locations.Results != nil {
		fmt.Println("Getting next locations...")
		printLocations(locations)
	}

	return nil
}

func GetPreviousLocations() error {
	var url string

	if locations.Previous == nil {
		fmt.Println("No previous locations.")
		return nil
	} else {
		url = locations.Previous.(string)
	}

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error getting next locations.")
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()

	if response.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d", response.StatusCode)
	}

	if err != nil {
		log.Fatalf("unmarshallLocationsreading response body.")
	}

	err = unmarshalLocations(body, &locations)

	if err != nil {
		log.Fatalf("Error unmarshalling locations response.")
	}

	if locations.Results != nil {
		fmt.Println("Getting previous locations...")
		printLocations(locations)
	}

	return nil
}
