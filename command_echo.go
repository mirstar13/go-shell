package main

import (
	"fmt"
	"strings"
)

func commandEcho(args ...string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}
