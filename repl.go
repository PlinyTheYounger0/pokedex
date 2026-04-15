package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(configPtr *config) {
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

		commandRegistry := getCommands()

		command, exists := commandRegistry[userCommand]
		if !exists {
			fmt.Println("Unknown Command")
		} else {
			err := command.callback(configPtr)
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
	next *string
	prev *string
}
