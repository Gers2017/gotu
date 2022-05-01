package typedefs

import (
	"fmt"
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

type CmdModule struct {
	Name        string
	Description string
	HelpText    string
	Actions     map[string]Action
}

func NewCmdModule(name, description, helptext string) CmdModule {
	return CmdModule{name, description, helptext, make(map[string]Action)}
}

type Action struct {
	Name     string
	HelpText string
}

type ActionFunc = func(args []string)

func (cmd *CmdModule) AddAction(k string, helpText string) {
	cmd.Actions[k] = Action{Name: k, HelpText: helpText}
}

func (cmd *CmdModule) PrintHelp() {
	fmt.Printf("[%s]\n  %s\n", cmd.Description, cmd.HelpText)
}
