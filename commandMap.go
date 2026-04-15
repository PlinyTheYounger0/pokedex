package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/PlinyTheYounger0/pokedex/internal/pokeapi"
)

func commandMap(configPtr *config) error {

	if configPtr.next == nil {
		baseURL := "https://pokeapi.co/api/v2/location-area"
		configPtr.next = &baseURL
	}

	locationAreas, err := getLocationAreas(*configPtr.next)
	if err != nil {
		return err
	}

	configPtr.next = locationAreas.Next
	configPtr.prev = locationAreas.Previous

	for _, la := range locationAreas.Results {
		fmt.Println(la.Name)
	}

	return nil
}

func commandMapb(configPtr *config) error {
	if configPtr.prev == nil {
		fmt.Println("You are on the first page")
		return nil
	}

	locationAreas, err := getLocationAreas(*configPtr.prev)
	if err != nil {
		return err
	}

	configPtr.next = locationAreas.Next
	configPtr.prev = locationAreas.Previous

	for _, la := range locationAreas.Results {
		fmt.Println(la.Name)
	}

	return nil
}

func getLocationAreas(url string) (*pokeapi.LocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving location areas: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading location areas: %w", err)
	}

	var locationAreas pokeapi.LocationAreas
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return nil, fmt.Errorf("Error unmarshaling location areas: %w", err)
	}

	return &locationAreas, nil
}
