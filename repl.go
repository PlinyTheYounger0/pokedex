package main

import (
	"strings"
)

func cleanInput(text string) []string {
	inputLower := strings.ToLower(text)
	words := strings.Fields(inputLower)
	return words
}
