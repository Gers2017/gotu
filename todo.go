package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
)

type Todo struct {
	Title    string
	Priority int
	Items    []string
}

func NewTodo(title string, priority int, items []string) Todo {
	return Todo{title, priority, items}
}

func (todo *Todo) AddItem(item string) {
	todo.Items = append(todo.Items, item)
}

func (todo *Todo) Print() {
	bangs := strings.Repeat("!", todo.Priority)

	contents := make([]string, 0)
	for _, item := range todo.Items {
		contents = append(contents, "  "+item)
	}

	t := color.New(color.FgCyan).Add(color.Bold)
	t.Printf("%s %s\n", todo.Title, bangs)
	c := color.New(color.FgYellow)
	c.Println(strings.Join(contents, "\n"))
}

func (todo *Todo) ToText() string {
	bangs := strings.Repeat("!", todo.Priority)
	contents := make([]string, 0)
	for _, item := range todo.Items {
		contents = append(contents, "  "+item)
	}

	return fmt.Sprintf("%s %s\n%s\n", todo.Title, bangs, strings.Join(contents, "\n"))
}

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
	priority := strings.Count(line[cut+1:], "!")
	return TitleTuple{title, priority}
}

func PrintTodos(todos []Todo) {
	if len(todos) == 0 {
		return
	}

	for _, todo := range todos {
		todo.Print()
	}
	fmt.Println() // add padding after printing
}

func TodosToText(todos []Todo) string {
	if len(todos) == 0 {
		return ""
	}

	todosText := make([]string, 0)
	for _, todo := range todos {
		todosText = append(todosText, todo.ToText())
	}

	return strings.Join(todosText, "\n")
}

func SortTodosByPriorityAsc(todos []Todo) {
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Priority < todos[j].Priority
	})
}

func GetTodos(filename string) []Todo {
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

	todos := make([]Todo, 0)

	for _, value := range titles {
		todo := NewTodo(value.title, value.priority, make([]string, 0))
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

	SortTodosByPriorityAsc(todos)
	return todos
}
