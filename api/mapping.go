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
	Previous any    `json:"previous"`

	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func UnmarshalLocations(body []byte) (LocationResponse, error) {
	locations:= LocationResponse{}
	err := json.Unmarshal(body, &locations)

	if err != nil {
		fmt.Println("Error unmarshalling locations.")
		return LocationResponse{}, err
	}

	return locations, nil
}


func GetNextLocations() error {
	response, err := http.Get("https://pokeapi.co/api/v2/location-area/")

	if err != nil {
		fmt.Println("Error getting next locations.")
	}

	body,err := io.ReadAll(response.Body)
	response.Body.Close()

	if response.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d", response.StatusCode)
	}
	
	if err != nil {
		log.Fatalf("Error reading response body.")
	}
	
	unmarshBody, err := UnmarshalLocations(body)
	
	if err != nil {
		return err
	}

	if(unmarshBody.Next != nil) {
		for _, location := range unmarshBody.Results {
			fmt.Println(location.Name)
	}
}

	return nil
}

func GetPreviousLocations() error {
	return nil
}