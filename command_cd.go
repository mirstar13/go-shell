package main

import (
	"fmt"
	"os"
)

func commandCd(args ...string) error {
	if len(args) < 1 {
		fmt.Println("cd: Invalid arguments")
		return nil
	}

	path := args[0]

	if path == "~" {
		path = os.Getenv("HOME")
	}

	if err := os.Chdir(path); err != nil {
		fmt.Printf("cd: %s: No such file or directory", path)
		fmt.Println()
	}

	return nil
}
