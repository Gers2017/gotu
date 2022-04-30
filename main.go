package main

import (
	"fmt"
	"os"
	"strings"
)

/*
  - get args
  - parse args
  - identify subcommand
  - populate data structure config
  - investigate how to pass the structure as a reference (avoid copies)
  - investigate, sort by -- and - priority in args
*/

type Command struct {
	Cmd         string
	Args        []string
	Len         int
	Description string
	HelpText    string
}

func NewCommand(_Args []string, _Description, _HelpText string) Command {
	_Len := len(_Args)
	_Cmd := _Args[0]
	_Args = _Args[1:]
	return Command{_Cmd, _Args, _Len, _Description, _HelpText}
}

func main() {

	args := os.Args[1:]
	// Use generics to map args to a lowercase slice

	fmt.Printf("%v %d\n", args, len(args))
	// config :=  Config{ Args: args, Len: len(args) }

	if len(args) < 1 {
		return
	}

	command := args[0] // maybe parse command func here

	switch command {
	case "get":
		getCmd := NewCommand(args, "Get all the configs", "Usage: get all simps")
		// fmt.Printf("GET Subcommand: %v\n", getCmd)
		HandleGet(&getCmd)
	case "set":
		fmt.Println("SET Subcommand!")

	case "help":
		fmt.Println("HELP command!")

	default:
		fmt.Println("Unknown subcommand")
	}

}

func HandleGet(get *Command) {
	subcommand := get.Args[0]
	subcommand = strings.ToLower(subcommand)

	switch subcommand {
	case "help":
		fmt.Printf("%s - %s\n", get.Description, get.HelpText)
	case "all":
		fmt.Println("Getting all items!")
	}
}
