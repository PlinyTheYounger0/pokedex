package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
			command.callback()
		}

	}
}
