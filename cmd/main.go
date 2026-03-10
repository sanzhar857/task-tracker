package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sanzhar857/task-tracker/internal/service"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("no command provided")
	}

	cmd := args[0]

	switch strings.ToLower(cmd) {
	case "help":
		if len(args) != 1 {
			return fmt.Errorf("usage: help")
		}
		service.Help()
		return nil
	case "add":
		if len(args) < 2 {
			return fmt.Errorf("usage: ./task-cli add \"<description>\"")
		}
		return service.AddTask(args[1])
	case "update":
		if len(args) < 3 {
			return fmt.Errorf("usage: ./task-cli update <id> \"<description>\"")
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("Invalid ID: %s", args[1])
		}

		return service.UpdateTask(id, args[2])
	case "delete":
		if len(args) < 2 {
			return fmt.Errorf("usage: ./task-cli delete <id>")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("Invalid ID: %s", args[1])
		}
		return service.Delete(id)
	}

	return nil
}
