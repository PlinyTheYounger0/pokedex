package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	locationsRes, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locationsRes.Next
	cfg.prevLocationsUrl = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsUrl == nil {
		return errors.New("You are on the first page")
	}

	locationsRes, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locationsRes.Next
	cfg.prevLocationsUrl = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
