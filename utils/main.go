package utils

import (
	"fmt"
)

func PrintUnknownAction(action string) {
	fmt.Printf("Unknown action: \"%s\"\n", action)
}

func GetArg(args []string, index int, defaultValue string) string {
	if len(args) <= index {
		return defaultValue
	}
	return args[index]
}

func GetArgsRange(args []string, start int) []string {
	if len(args) <= start {
		return make([]string, 0)
	}
	return args[start:]
}

func PrintHelp(title, helpText string) {
	fmt.Printf("[%s]\n  %s\n", title, helpText)
}
