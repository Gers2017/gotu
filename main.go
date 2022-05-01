package main

import (
	"fmt"
	"gotu/modules"
	"log"
	"os"

	. "gotu/typedefs"
	. "gotu/utils"
)

func main() {
	args := os.Args[1:]
	gotu := NewCmdModule("gotu", "Manage your gotus", "Gotu core commands: get, help")

	gotu.AddAction("get", "Get the todos")

	if len(args) < 1 {
		gotu.PrintHelp()
		return
	}

	action := GetArg(args, 0, "")  // get
	flags := GetArgsRange(args, 1) // [all, --title, --primary]

	fmt.Println("ACTION:", action, "ARGS:", flags)

	switch action {
	case "get":
		if len(flags) < 1 {
			log.Fatalln("Missing Parameters")
		}
		modules.HandleGetCmd(flags, gotu.Actions["get"])
	default:
		gotu.PrintHelp()
	}
}
