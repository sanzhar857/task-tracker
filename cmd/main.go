package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("The required format to continue working: ./task-tracker [command] [arguments]")
		fmt.Println("Or enter a \"help\" command to view a list of available commands.")
		return
	}

	command := args[0]
	switch strings.TrimSpace(command) {
	case "add":
	case "update":
	case "delete":
	case "list":
	default:
		fmt.Println("Unknown command.")
	}
}
