package main

import (
	"fmt"
	"os"
)

func commandPwd(args ...string) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not retrieve current working directory err: %w", err)
	}
	fmt.Println(wd)

	return nil
}
