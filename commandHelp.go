package main

import (
	"fmt"
)

func commandHelp(command string) {
	fmt.Println("Poké CLI is a command-line interface for retrieving information from the Poké API.")

	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  pokedex-cli <command>")

	fmt.Println("")
	fmt.Println("Available commands:")
	fmt.Println("  help    Display this help message")
	fmt.Println("")
}
