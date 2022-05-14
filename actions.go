package main

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintAllTodos(conf *Config) {
	blue := color.New(color.FgBlue).Add(color.Bold)
	blue.Println("Get all the todos")
	todos := GetTodos(conf.TodoFile)
	PrintTodos(todos)
}

func PrintPrimaryTodo(conf *Config) {
	todos := GetTodos(conf.TodoFile)
	SortTodosByPriority(todos)

	if len(todos) < 1 {
		fmt.Println("Empty todos!")
	}

	todos[0].Print()
}

func PrintTodoByTitle(title string, conf *Config) {
	if title == "" {
		fmt.Println("Empty title!")
		return
	}

	todos := GetTodos(conf.TodoFile)

	found := false
	for _, todo := range todos {
		if clearTitle(todo.Title) == clearTitle(title) {
			todo.Print()
			found = true
			break
		}
	}

	if !found {
		fmt.Println("No todo with title:", title, "on todos!")
	}
}
