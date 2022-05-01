package utils

import (
	"fmt"
	"gotu/typedefs"
	"io/ioutil"
	"os"
	"strings"
)

func isTodoTitle(line string) bool {
	return strings.HasPrefix(line, "[") && strings.Contains(line, "]")
}

type TitleTuple struct {
	title    string
	priority int
}

func parseTitle(line string) TitleTuple {
	cut := strings.Index(line, "]")
	title := line[:cut+1]
	priority := strings.Count(line[:cut], "!")
	return TitleTuple{title, priority}
}

func TodosToText(todos []typedefs.Todo) string {
	if len(todos) == 0 {
		return ""
	}
	todosText := make([]string, 0)
	for _, todo := range todos {
		todosText = append(todosText, todo.ToText())
	}
	return strings.Join(todosText, "\n")
}

func GetTodos(filename string) []typedefs.Todo {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error uknown file: ", filename)
		os.Exit(1)
	}

	text := string(content)
	lines := make([]string, 0)
	titles := make([]TitleTuple, 0)

	for _, line := range strings.Split(text, "\n") {
		if len(line) > 0 {
			lines = append(lines, strings.Trim(line, " "))
			if isTodoTitle(line) {
				titles = append(titles, parseTitle(line))
			}
		}
	}

	todos := make([]typedefs.Todo, 0)

	for _, value := range titles {
		todo := typedefs.NewTodo(value.title, value.priority, make([]string, 0))
		todos = append(todos, todo)
	}

	index := 0

	for _, line := range lines[1:] {

		if isTodoTitle(line) {
			index += 1
		} else {
			todos[index].AddItem(strings.Trim(line, " "))
		}
	}

	// sort todos...
	return todos
}
