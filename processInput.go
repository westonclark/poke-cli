package main

import (
	"fmt"
	"os"
	"strings"
)

func processInput(text string) {
	words := cleanInput(text)

	if len(words) == 0 {
		os.Exit(0)
	}

	command := words[0]
	proccessCommand(command)
}

func proccessCommand(command string) {
	switch command {
	case "help":
		commandHelp(command)
	default:
		fmt.Println("Finding Pokémon...", command)
	}
}

func cleanInput(input string) []string {
	lowerCase := strings.ToLower(input)
	words := strings.Fields(lowerCase)

	return words
}
