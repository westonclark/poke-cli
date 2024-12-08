package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	for {
		fmt.Print("Enter a pokémon's name to lean more: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()

		words := cleanInput(text)

		if len(words) == 0 {
			continue
		}

		command := words[0]

		switch command {
		case "help":
			fmt.Println("Poké CLI is a command-line interface for retrieving information from the Poké API.")

			fmt.Println("")
			fmt.Println("Usage:")
			fmt.Println("  pokedex-cli <command>")

			fmt.Println("")
			fmt.Println("Available commands:")
			fmt.Println("  help    Display this help message")
			fmt.Println("  exit    Exit the program")
			fmt.Println("")

		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
		}

	}
}

func cleanInput(input string) []string {
	lowerCase := strings.ToLower(input)
	words := strings.Fields(lowerCase)

	return words
}
