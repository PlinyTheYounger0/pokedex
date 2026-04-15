package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commandRegistry := getCommands()

	for i := range commandRegistry {
		command := commandRegistry[i]

		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
