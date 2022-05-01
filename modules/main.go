package modules

import (
	"fmt"
	"gotu/typedefs"
	"gotu/utils"
)

type Action = typedefs.Action

func HandleGetCmd(flags []string, action Action) {
	flag := utils.GetArg(flags, 0, "")

	switch flag {
	case "all":
		getAllTodos(flags)
	case "primary":
		getPrimaryTodo(flags)
	case "title":
		getTodoByTitle(flags)
	default:
		utils.PrintHelp("Get", action.HelpText)
	}
}

func getAllTodos(args []string) {
	fmt.Println("Get all the todos")
	todos := utils.GetTodos("test.tudu")
	fmt.Println(utils.TodosToText(todos))
}

func getPrimaryTodo(args []string) {
	fmt.Println("Get primary todo")
}

func getTodoByTitle(args []string) {
	title := utils.GetArg(args, 1, "")
	if title == "" {
		fmt.Println("Missing title parameter!")
		return
	}
	fmt.Printf("Get todo by title %s\n", title)
}
