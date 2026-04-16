package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PlinyTheYounger0/pokedex/internal/pokeapi"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanned := scanner.Scan()
		if scanned == false {
			fmt.Printf("Error scanning: %v", scanner.Err())
		}

		text := scanner.Text()
		input := cleanInput(text)
		userCommand := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		commandRegistry := getCommands()

		command, exists := commandRegistry[userCommand]
		if !exists {
			fmt.Println("Unknown Command")
		} else {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}

func cleanInput(text string) []string {
	inputLower := strings.ToLower(text)
	words := strings.Fields(inputLower)
	return words
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsUrl *string
	prevLocationsUrl *string
	pokedex          map[string]pokeapi.Pokemon
}
