package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func trimWhitespace(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func parseFlag(f string) string {
	return strings.ReplaceAll(f, "-", "")
}

func IsFlag(f string) bool {
	return strings.HasPrefix(f, "-") && len(f) > 1
}

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

	flags := conf.Args[2:]

	for _, f := range flags {
		if IsFlag(f) {
			key := parseFlag(f)
			key = trimWhitespace(key)
			conf.Flags[key] = true
		}
	}

	conf.Action = action
	return conf
}

func (conf *Config) GetArg(index int) (string, error) {
	args := conf.Args
	if len(args) <= index {
		return "", errors.New(fmt.Sprintf("Trying to access args[%d], args len: %d", index, len(args)))
	}
	return args[index], nil
}

func (conf *Config) HasFlag(flagKey string) bool {
	_, exists := conf.Flags[flagKey]
	return exists
}

func (conf *Config) AddCommand(key string, _command Command) {
	conf.CommandMap[key] = _command
}

type Command struct {
	Name     string
	Flags    map[string]Flag
	FlagSet  []string
	HelpText string
}

func NewCommand(_name, _help string) Command {
	return Command{_name, make(map[string]Flag), make([]string, 0), _help}
}

func (cmd *Command) GetVariantsOf(key string) ([]string, error) {
	f, ok := cmd.Flags[key]
	if !ok {
		return make([]string, 0), errors.New(fmt.Sprintf("No such a flag: %s", key))
	}

	return f.Variants, nil
}

func (cmd *Command) AddFlag(f Flag) {
	cmd.Flags[f.Name] = f
	cmd.FlagSet = append(cmd.FlagSet, f.Variants...)
}

type Flag struct {
	Name         string
	Variants     []string
	DefaultValue string
	HelpText     string
}

func NewFlag(name string, variants string, defaultValue string, helpText string) Flag {
	return Flag{name, strings.Split(variants, " "), defaultValue, helpText}
}

func PrintUnknown(value string) {
	fmt.Printf("Unknown command: \"%s\"\n", value)
}

func Matches(value string, variants []string) bool {
	matches := false
	for _, variant := range variants {
		if value == variant {
			return true
		}
	}
	return matches
}

func MatchesRange(args []string, variants []string) bool {
	variantStr := strings.Join(variants, " ")

	for _, arg := range args {
		if strings.Contains(variantStr, arg) {
			return true
		}
	}

	return false
}

func (conf *Config) HasFlagInVariant(variants []string) bool {
	for k := range conf.Flags {
		for _, v := range variants {
			if parseFlag(v) == k {
				return true
			}
		}
	}

	return false
}

const (
	ROOT_HELP        = "Gotu core commands: [get, rm, add].\nUsage: gotu get --all"
	GET_ACTION_HELP  = "Get action help: tudu get [--all, --title, --primary]"
	GET_ALL_HELP     = "Get all: tudu get [--all, -A, --todos]"
	GET_TITLE_HELP   = "Get title: tudu get [--title, -T] <your-title>"
	GET_PRIMARY_HELP = "Get primary: tudu get [--primary, -P]"
)

func main() {
	// ind:  0   1     2      3
	// 		[todo get --title "clocks"]
	// len:  1   2     3      4
	config := NewConfig(os.Args)

	for k, v := range config.Flags {
		fmt.Println(k, ":", v)
	}

	getCommand := NewCommand("get", GET_ACTION_HELP)
	getCommand.AddFlag(NewFlag("all", "--all -A --todos", "false", GET_ALL_HELP))
	getCommand.AddFlag(NewFlag("title", "--title -T", "", GET_TITLE_HELP))
	getCommand.AddFlag(NewFlag("help", "--help -h", "", GET_ACTION_HELP))

	config.AddCommand("get", getCommand)

	HELP_VARIANTS := []string{"help", "--help", "-h"}
	isHelp := config.HasFlagInVariant(HELP_VARIANTS)

	// Match value, on fail print error message or help message
	if config.Action == "get" {
		if config.HasFlag("all") || config.HasFlag("A") || config.HasFlag("todos") {
			if isHelp {
				fmt.Println(getCommand.Flags["all"].HelpText)
				return
			}

			fmt.Println("Printing all values...")

		} else if config.HasFlag("title") || config.HasFlag("T") {

			if isHelp {
				fmt.Println(getCommand.Flags["title"].HelpText)
				return
			}

			title, err := config.GetArg(3)

			if err != nil {
				fmt.Println(getCommand.Flags["title"].HelpText)
				return
			}

			fmt.Printf("Printing todos by title: \"%s\"...\n", title)
		}

		if isHelp {
			println(GET_ACTION_HELP)
			return
		}
	} else {
		fmt.Println(GET_ACTION_HELP)
	}
}
