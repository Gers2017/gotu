package main

import (
	"fmt"
	"os"

	"github.com/Gers2017/flago"
)

func PrintUnknown(value string) {
	fmt.Printf("Unknown command: \"%s\"\n", value)
}

const (
	ROOT_HELP        = "Gotu core commands: [get, rm, add].\nUsage: gotu get --all"
	GET_ACTION_HELP  = "Get action help: tudu get [--all, --title, --primary]"
	GET_ALL_HELP     = "Get all: tudu get [--all, -A, --todos]"
	GET_TITLE_HELP   = "Get title: tudu get [--title, -T] <your-title>"
	GET_PRIMARY_HELP = "Get primary: tudu get [--primary, -P]"
)

func main() {
	config := NewConfig(os.Args, "test.tudu")
	get := flago.NewFlagSet("get")
	get.Bool("all", false)
	get.Bool("primary", false)
	get.Str("title", "<none>")
	get.Bool("help", false)

	get.ParseFlags(config.Args)

	switch config.Action {
	case "get":
		hasHelp := get.HasFlag("help")
		hasAll := get.HasFlag("all")
		hasPriority := get.HasFlag("primary")
		hasTitle := get.HasFlag("title")

		if hasAll {
			if hasHelp {
				fmt.Println(GET_ALL_HELP)
				return
			}

			PrintAllTodos(&config)

		} else if hasPriority {
			if hasHelp {
				fmt.Println(GET_PRIMARY_HELP)
				return
			}

			PrintPrimaryTodo(&config)

		} else if hasTitle {
			if hasHelp {
				fmt.Println(GET_TITLE_HELP)
				return
			}

			title := get.GetStr("title")
			PrintTodoByTitle(title, &config)

		} else {
			fmt.Println(ROOT_HELP)
		}

	default:
		fmt.Println(ROOT_HELP)
	}
}
