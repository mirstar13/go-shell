package main

import (
	"fmt"
	"os"
)

func commandLs(args ...string) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not retrieve current working directory")
	}

	files, err := os.ReadDir(wd)
	if err != nil {
		return fmt.Errorf("could not read current working directory")
	}

	fmt.Println("   Mode       Type       Length Name")
	fmt.Println()

	for _, file := range files {
		fileType := ""

		info, err := file.Info()
		if err != nil {
			return err
		}

		if file.IsDir() {
			fileType = "Directory"
		} else {
			fileType = "File"
		}

		fmt.Printf("%-10s %10s %9d %-10s\n", info.Mode(), fileType, info.Size(), file.Name())
	}

	return nil
}
