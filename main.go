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

	conf.Action = action
	return conf
}

func (conf *Config) ParseFlags() {
	fmt.Println("Parsing flags...")
	flags := conf.Args[1:]

	for _, f := range flags {
		if IsFlag(f) {
			userFlag := parseFlag(f)
			userFlag = trimWhitespace(userFlag)

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

func (conf *Config) HasFlag(flagKey string) bool {
	_, exists := conf.Flags[flagKey]
	return exists
}

func (conf *Config) AddCommand(key string, _command Command) {
	conf.CommandMap[key] = _command
}

type Command struct {
	Name         string
	Flags        map[string]Flag
	FlagVariants map[string][]string
	HelpText     string
}

func NewCommand(_name, _help string) Command {
	return Command{_name, make(map[string]Flag), make(map[string][]string), _help}
}

func (cmd *Command) PrintHelp() {
	fmt.Println(cmd.HelpText)
}

func (cmd *Command) GetVariantsOf(flagName string) ([]string, error) {
	v, ok := cmd.FlagVariants[flagName]
	if !ok {
		return make([]string, 0), errors.New(fmt.Sprintf("Unknown flag \"%s\" \n", flagName))
	}
	return v, nil
}

func (cmd *Command) HasVariant(val string) (bool, string) {
	for key, variants := range cmd.FlagVariants {
		for _, v := range variants {
			v = trimWhitespace(v)
			v = parseFlag(v)

			if v == val {
				return true, key
			}
		}
	}

	return false, ""
}

func (cmd *Command) AddFlag(f Flag) {
	cmd.Flags[f.Name] = f
	cmd.FlagVariants[f.Name] = f.Variants
}

type Flag struct {
	Name         string
	Variants     []string
	DefaultValue any
	HelpText     string
}

func NewFlag(name string, variants string, defaultValue any, helpText string) Flag {
	return Flag{name, strings.Split(variants, " "), defaultValue, helpText}
}

func (flag *Flag) PrintHelp() {
	fmt.Println(flag.HelpText)
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

func PrintHelpFor(key string) {

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

	get := NewCommand("get", GET_ACTION_HELP)
	get.AddFlag(NewFlag("all", "--all -A --todos", false, GET_ALL_HELP))
	get.AddFlag(NewFlag("primary", "--primary -P", false, GET_PRIMARY_HELP))
	get.AddFlag(NewFlag("title", "--title -T", "", GET_TITLE_HELP))
	get.AddFlag(NewFlag("help", "--help -h", false, GET_ACTION_HELP))
	config.AddCommand("get", get)

	config.ParseFlags()

	isHelp := config.HasFlag("help")

	// Match value, on fail print error message or help message
	if config.Action == "get" {

		if config.HasFlag("all") {
			if isHelp {
				fmt.Println(get.Flags["all"].HelpText)
				return
			}

			fmt.Println("Printing all values...")

		} else if config.HasFlag("primary") {

			fmt.Println("Printing priamry value...")

		} else if config.HasFlag("title") {

			if isHelp {
				fmt.Println(get.Flags["title"].HelpText)
				return
			}

			title, err := config.GetArg(3) // This doesn't work anymore

			if err != nil {
				fmt.Println(get.Flags["title"].HelpText)
				return
			}

			fmt.Printf("Printing todos by title: \"%s\"...\n", title)
		}

		if isHelp {
			get.PrintHelp()
		}

	} else {
		get.PrintHelp()
	}
}
