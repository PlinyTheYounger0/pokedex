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

		fmt.Printf("Your first word was: %s\n", input[0])
	}
}
