package main

import (
	"errors"
	"fmt"
)

func GetArg(args []string, index int) (string, error) {
	if len(args) <= index {
		return "", errors.New(fmt.Sprintf("Index out of range. Trying to access args[%d]", index))
	}
	return args[index], nil
}

func PrintHelp(title, helpText string) {
	fmt.Printf("[%s]\n  %s\n", title, helpText)
}
