package main

import (
	"fmt"

	"github.com/fatih/color"
)

func getAllTodos() {
	blue := color.New(color.FgBlue).Add(color.Bold)
	blue.Println("Get all the todos")
	todos := GetTodos("test.tudu")
	PrintTodos(todos)
}

func getPrimaryTodo() {
	fmt.Println("Get primary todo...")
}

func getTodoByTitle(title string) {
	if title == "" {
		fmt.Println("Missing title parameter!")
		return
	}
	fmt.Printf("Get todos by title: \"%s\"...\n", title)
}
