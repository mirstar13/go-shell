package main

import (
	"fmt"
)

func commandType(args ...string) error {
	if len(args) == 0 {
		fmt.Println("type: Invalid arguments")
		return nil
	}

	commands := getCommands()

	if command, exists := commands[args[0]]; exists {
		fmt.Printf("%s is a shell builtin", command.Name)
		fmt.Println()
	} else {
		if file, err := handleEnviromentVariablePaths(args[0]); err == nil {
			fmt.Printf("%s is %s", args[0], file)
			fmt.Println()
		} else {
			fmt.Printf("%s not found", args[0])
			fmt.Println()
		}
	}

	return nil
}
