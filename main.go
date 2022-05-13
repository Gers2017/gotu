package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Args       []string
	Action     string
	Flags      map[string]bool
	CommandMap map[string]Command
}

func NewConfig(_args []string) Config {
	conf := Config{_args, "", make(map[string]bool), make(map[string]Command)}

	action, err := conf.GetArg(1)

	if err != nil {
		fmt.Println(ROOT_HELP)
		os.Exit(0)
	}

	conf.Action = action
	return conf
}

func (conf *Config) ParseFlags() {
	fmt.Println("Parsing flags...")
	flags := conf.Args[1:]

	for _, f := range flags {
		if IsFlag(f) {
			userFlag := sanitize(f)

			for _, cmd := range conf.CommandMap {
				hasVariant, flagId := cmd.HasVariant(userFlag)
				if hasVariant {
					conf.Flags[flagId] = true
				}
			}

		}
	}

	for k, v := range conf.Flags {
		fmt.Println(k, "->", v)
	}
}

func (conf *Config) GetArg(index int) (string, error) {
	args := conf.Args
	if len(args) <= index {
		return "", errors.New(fmt.Sprintf("Trying to access args[%d], args len: %d", index, len(args)))
	}
	return args[index], nil
}

func (conf *Config) IndexOf(arg string) int {
	for i, v := range conf.Args {
		if arg == sanitize(v) {
			return i
		}
	}
	return -1
}

func (conf *Config) HasFlag(flagKey string) bool {
	_, exists := conf.Flags[flagKey]
	return exists
}

func (conf *Config) AddCommand(key string, _command Command) {
	conf.CommandMap[key] = _command
}

func sanitize(s string) string {
	s = parseFlag(s)
	return trimWhitespace(s)
}

func trimWhitespace(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func parseFlag(f string) string {
	return strings.ReplaceAll(f, "-", "")
}

func IsFlag(f string) bool {
	return strings.HasPrefix(f, "-") && len(f) > 1
}

type Command struct {
	Name     string
	Flags    map[string]Flag
	HelpText string
}

func NewCommand(_name, _help string) Command {
	return Command{_name, make(map[string]Flag), _help}
}

func (cmd *Command) PrintHelp() {
	fmt.Println(cmd.HelpText)
}

func (cmd *Command) HasVariant(val string) (bool, string) {
	for key, flag := range cmd.Flags {
		for _, v := range flag.Variants {
			v = sanitize(v)

			if v == val {
				return true, key
			}
		}
	}

	return false, ""
}

func (cmd *Command) AddFlag(name string, variants string, defaultValue any, helpText string) {
	f := NewFlag(name, variants, defaultValue, helpText)
	cmd.Flags[f.Name] = f
}

type Flag struct {
	Name         string
	Variants     []string
	DefaultValue any
	DataType     string
	IsRequired   bool
	HelpText     string
}

func NewFlag(name string, variants string, defaultValue any, helpText string) Flag {
	dataType := fmt.Sprint(reflect.TypeOf(defaultValue))
	isRequired := dataType != "bool"
	return Flag{name, strings.Split(variants, " "), defaultValue, dataType, isRequired, helpText}
}

func (flag *Flag) PrintHelp() {
	fmt.Println(flag.HelpText)
}

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
	// ind:  0   1     2      3       4
	// 		[todo get --title <title> --help]
	// len:  1   2     3      4       5
	config := NewConfig(os.Args)

	get := NewCommand("get", GET_ACTION_HELP)
	get.AddFlag("all", "--all -A --todos", false, GET_ALL_HELP)
	get.AddFlag("primary", "--primary -P", false, GET_PRIMARY_HELP)
	get.AddFlag("title", "--title -T", "", GET_TITLE_HELP)
	get.AddFlag("help", "--help -h", false, GET_ACTION_HELP)
	config.AddCommand("get", get)

	config.ParseFlags()

	isHelp := config.HasFlag("help")

	if config.Action == "get" {

		if config.HasFlag("all") {
			if isHelp {
				fmt.Println(get.Flags["all"].HelpText)
				return
			}

			getAllTodos()

		} else if config.HasFlag("primary") {
			if isHelp {
				fmt.Println(get.Flags["primary"].HelpText)
				return
			}

			getPrimaryTodo()

		} else if config.HasFlag("title") {
			if isHelp {
				fmt.Println(get.Flags["title"].HelpText)
				return
			}

			tIndex := config.IndexOf("title")
			if tIndex == -1 {
				fmt.Println(get.Flags["title"].HelpText)
				return
			}

			title, err := config.GetArg(tIndex + 1) // Search for the <title> value after the flag

			if err != nil {
				fmt.Println(get.Flags["title"].HelpText)
				return
			}

			getTodoByTitle(title)
		}

		if isHelp {
			get.PrintHelp()
		}
	} else {
		get.PrintHelp()
	}
}
