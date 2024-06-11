package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Command struct {
	Name     string
	Type     string
	Callback func(args ...string) error
}

func printPrompt() {
	fmt.Fprint(os.Stdout, "$ ")
}

func cleanInput(s string) string {
	result := strings.TrimSpace(s)
	return result
}

func startRepl() error {
	commands := getCommands()

	reader := bufio.NewScanner(os.Stdin)
	printPrompt()

	for reader.Scan() {
		text := cleanInput(reader.Text())
		spltText := strings.Split(text, " ")
		if command, exists := commands[strings.ToLower(spltText[0])]; exists {
			err := command.Callback(spltText[1:]...)
			if err != nil {
				return fmt.Errorf("error running command err: %w", err)
			}
		} else {
			if file, err := handleEnviromentVariablePaths(spltText[0]); err == nil {
				cmd := exec.Command(file, spltText[1:]...)
				cmd.Stdout = os.Stdout

				if err := cmd.Run(); err != nil {
					return fmt.Errorf("error running external command err: %w", err)
				}
			} else {
				fmt.Printf("%s: command not found\n", spltText[0])
			}
		}

		printPrompt()
	}

	return nil
}

func getCommands() map[string]Command {
	return map[string]Command{
		"exit": {
			Name:     "exit",
			Type:     "builtin",
			Callback: commandExit,
		},
		"echo": {
			Name:     "echo",
			Type:     "builtin",
			Callback: commandEcho,
		},
		"type": {
			Name:     "type",
			Type:     "builtin",
			Callback: commandType,
		},
		"pwd": {
			Name:     "pwd",
			Type:     "builtin",
			Callback: commandPwd,
		},
		"cd": {
			Name:     "cd",
			Type:     "builtin",
			Callback: commandCd,
		},
		"ls": {
			Name:     "ls",
			Type:     "builtin",
			Callback: commandLs,
		},
	}
}

func handleEnviromentVariablePaths(base string) (string, error) {
	paths := strings.Split(os.Getenv("PATH"), ":")

	var filePath string
	for _, dir := range paths {
		filePath = filepath.Join(dir, base)

		if _, err := os.Stat(filePath); err == nil {
			return filePath, nil
		}
	}

	return "", fmt.Errorf("file not found in specified paths")
}
