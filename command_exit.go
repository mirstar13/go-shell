package main

import (
	"fmt"
	"os"
	"strconv"
)

func commandExit(args ...string) error {
	if len(args) == 0 {
		args = append(args, "0")
	}

	exitSatus, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid exit code/status argument err: %w", err)
	}

	os.Exit(exitSatus)
	return nil
}
